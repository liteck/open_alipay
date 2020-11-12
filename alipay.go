package open_alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

/**
接口类.
实现时需要满足下面条件
*/
type api interface {
	// 支付宝的Method值
	getMethod() (method string)
	// 请求参数.指定请求参数属于哪里的
	getReq() (biz interface{}, form interface{})
}

type baseApi struct {
	Input interface{} // 请求参数
}

/**
支付宝
*/
type AlipayClient struct {
	AppId      string `validate:"required"` //开放平台建立的appId
	UserId     string `validate:"required"` //登录账户的 UserId
	PublicRSA  []byte `validate:"required"` //支付宝公钥.接口验签使用
	PrivateRSA []byte `validate:"required"` //支付宝私钥.接口签名使用.(注:非 java 的 pcsk8.)
}

func (a *AlipayClient) Execute(_api api, token string, ptrResp interface{}) (err error) {
	// check
	if err = valid(a); err != nil {
		log.Println(err.Error())
		return
	}

	//公共参数
	params := a.initParams(_api.getMethod(), token)

	//业务参数
	bizContent, formParams := _api.getReq()
	// biz类参数可不是这样的
	if bizContent != nil {
		var bizData []byte
		if bizData, err = json.Marshal(&bizContent); err != nil {
			log.Println(err.Error())
			return
		}
		params.Set("biz_content", string(bizData))
	}

	// 其它表单类参数是拼接
	if formParams != nil {
		oMap := jsonToMap(formParams)
		for k, v := range oMap {
			params.Set(k, fmt.Sprintf("%v", v))
		}
	}

	// 签名
	var sign string
	if sign, err = a.sign(params); err != nil {
		log.Println(err.Error())
		return
	} else {
		params.Set("sign", sign)
	}

	// 请求
	var res *http.Response
	if res, err = http.PostForm("https://openapi.alipay.com/gateway.do", params); err != nil {
		log.Println(err.Error())
		return
	}
	if res.Body == nil {
		err = errors.New("response body is null")
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	// 响应
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		log.Println(err.Error())
		return
	}
	respStr := string(body)
	log.Println("response data======" + respStr)

	// 最初的数据..应该是 k v sign
	r := map[string]interface{}{}
	if err = json.Unmarshal(body, &r); err != nil {
		log.Println(err.Error())
		return
	}

	var data string
	//把 method 转换为 key
	respKey := strings.Replace(_api.getMethod(), ".", "_", -1) + "_response"
	// 有时候返回这个错误,替换为
	respStr = strings.Replace(respStr, "error_response", respKey, -1)
	//原始签名
	if v, pass := a.verifySign(respStr, fmt.Sprintf("%v", r["sign"]), respKey); !pass {
		log.Println("verify failed")
	} else {
		data = v
	}

	var newBody []byte
	if newBody, err = gbkToUtf8([]byte(data)); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(newBody))

	if ptrResp == nil {
		ptrResp = map[string]interface{}{}
	}

	// 转义 GBK
	if err = json.Unmarshal(newBody, ptrResp); err != nil {
		log.Println(err.Error())
		return
	}

	return
}

/**
初始化
支付宝公共参数
*/
func (a *AlipayClient) initParams(method, token string) url.Values {
	p := url.Values{}

	p.Set("app_id", a.AppId)
	p.Set("method", method)
	p.Set("format", "JSON")
	p.Set("charset", "utf-8")
	p.Set("sign_type", "RSA")
	p.Set("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	p.Set("version", "1.0")
	p.Set("app_auth_token", token)

	return p
}

func (a *AlipayClient) sign(params url.Values) (sign string, err error) {
	//对key进行升序排序.
	keys := make([]string, 0)
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str string
	//对key=value的键值对用&连接起来，略过空值
	for _, k := range keys {
		value := params.Get(k)
		if value != "" {
			str = str + k + "=" + value + "&"
		}
	}
	str = strings.TrimRight(str, "&")

	log.Println("signed data====" + str)

	//签名
	block, _ := pem.Decode(a.PrivateRSA)
	if block == nil {
		err = errors.New("私钥错误")
		return
	}

	var private *rsa.PrivateKey
	if private, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return
	}

	_t := crypto.SHA1.New()
	_t.Write([]byte(str))
	digest := _t.Sum(nil)
	var data []byte
	if data, err = rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, digest); err != nil {
		return "", err
	}
	sign = base64.StdEncoding.EncodeToString(data)
	return
}

func (a *AlipayClient) verifySign(s, sign, respKey string) (signData string, pass bool) {
	idx := strings.Index(s, ",\"sign\"")
	if idx == -1 {
		return
	}
	signData = s[4+len(respKey) : idx]
	log.Println("verify signed data====" + signData)

	block, _ := pem.Decode(a.PublicRSA)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse RSA public key: %s\n", err)
		return
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	t := sha1.New()
	_, _ = io.WriteString(t, signData)
	digest := t.Sum(nil)
	data, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println("DecodeString sig error, reason: ", err)
		return
	}
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
		return
	}

	pass = true
	return
}
