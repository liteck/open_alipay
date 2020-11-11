package open_alipay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"time"
)

/**
支付宝
*/
type AlipayClient struct {
	AppId      string `validate:"required"` //开放平台建立的appId
	UserId     string `validate:"required"` //登录账户的 UserId
	PublicRSA  []byte `validate:"required"` //支付宝公钥.接口验签使用
	PrivateRSA []byte `validate:"required"` //支付宝私钥.接口签名使用.(注:非 java 的 pcsk8.)
}

func (a *AlipayClient) Execute(_api api, token string) (err error) {
	// check
	if err = valid(a); err != nil {
		log.Println(err.Error())
		return
	}

	//公共参数
	params := a.commonParams(_api.method(), token)
	if err = params.Valid(); err != nil {
		log.Println(err.Error())
		return
	}

	cMap := a.jsonToMap(params)

	//业务参数
	if b, o, e := _api.params(); e != nil {
		err = e
		log.Println(err.Error())
		return
	} else {
		// biz类参数可不是这样的
		if b != nil {
			//bMap := a.toMap(b)
			//for k, v := range bMap {
			//	cMap[k] = v
			//}
		}
		// 其它表单类参数是拼接
		if o != nil {
			oMap := a.jsonToMap(o)
			for k, v := range oMap {
				cMap[k] = v
			}
		}
	}

	log.Println(String(cMap))

	//字符串
	preSign := a.mapToString(cMap)
	log.Println(preSign)
	// 签名
	var sign string
	if sign, err = a.sign(preSign); err != nil {
		log.Println(err.Error())
		return
	} else {
		cMap["sign"] = sign
	}
	log.Println(String(cMap))

	// 请求
	var res *http.Response
	if res, err = http.PostForm("https://openapi.alipay.com/gateway.do", a.convertMap2UValue(cMap)); err != nil {
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

	//验签之类的..回头再说

	log.Println(string(body))

	var newBody []byte
	if newBody, err = GbkToUtf8(body); err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(newBody))

	// 转义 GBK
	if err = json.Unmarshal(newBody, _api.response()); err != nil {
		log.Println(err.Error())
		return
	}

	return
}

func (a *AlipayClient) commonParams(method, token string) (p CommonParams) {

	p.AppId = a.AppId
	p.Method = method
	p.Format = "JSON"
	p.Charset = "utf-8"
	p.SignType = "RSA"
	//p.Sign = ""
	p.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	p.Version = "1.0"
	p.AppAuthToken = token

	return
}

func (a *AlipayClient) getRespKey(method string) (key string) {
	key = strings.Replace(method, ".", "_", -1)
	return
}

func (a *AlipayClient) convertMap2UValue(__map map[string]interface{}) url.Values {
	ret := url.Values{}
	for k, v := range __map {
		ret[k] = []string{fmt.Sprintf("%v", v)}
	}
	return ret
}

func (a *AlipayClient) jsonToMap(params interface{}) map[string]interface{} {
	t := reflect.TypeOf(params)
	v := reflect.ValueOf(params)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		tag := t.Field(i).Tag.Get("json")
		if tag != "" {
			if strings.Contains(tag, ",") {
				ps := strings.Split(tag, ",")
				key = ps[0]
			} else {
				key = tag
			}
		}
		data[key] = value
	}
	return data
}

func (a *AlipayClient) mapToString(m map[string]interface{}) string {
	//对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", m[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if len(signStrings) == 0 {
		return ""
	} else {
		signStrings = signStrings[:len(signStrings)-1]
	}
	return signStrings
}

func (a *AlipayClient) sign(c string) (sign string, err error) {
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
	_t.Write([]byte(c))
	digest := _t.Sum(nil)
	var data []byte
	if data, err = rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, digest); err != nil {
		return "", err
	}
	sign = base64.StdEncoding.EncodeToString(data)
	return
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
