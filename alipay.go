package open_alipay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
接口类.
实现时需要满足下面条件
*/
type apiI interface {
	getMethod() (method string)
	getAppAuthToken() string
	getReq() interface{}
}

type api struct {
	appAuthToken string
	input        interface{} // Biz 类请求参数
}

func (a *api) SetParams(biz interface{}) {
	a.input = biz
}

func (a *api) SetAuthToken(token string) {
	a.appAuthToken = token
}

/**
支付宝
*/
type AlipayClient struct {
	AppId      string `validate:"required"` //开放平台建立的appId
	UserId     string `validate:"required"` //登录账户的 UserId
	PublicRSA  []byte `validate:"required"` //支付宝公钥.接口验签使用
	PrivateRSA []byte `validate:"required"` //支付宝私钥.接口签名使用.(注:非 java 的 pcsk8.)
	timeout    int                          // 超时时间: 3
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
	if token != "" {
		p.Set("app_auth_token", token)
	}

	return p
}

func (a *AlipayClient) Timeout(second int) *AlipayClient {
	a.timeout = second
	return a
}

/**
用于Biz类型参数接口
*/
func (a *AlipayClient) Execute(_api apiI, ptrResp interface{}) (err error) {
	// check
	if err = valid(a); err != nil {
		Trace.Println(err.Error())
		return
	}

	//系统参数
	params := a.initParams(_api.getMethod(), _api.getAppAuthToken())

	//业务参数
	bizContent := _api.getReq()
	// biz类参数可不是这样的
	if bizContent != nil {
		var bizData, bizData2 []byte
		if bizData, err = json.Marshal(&bizContent); err != nil {
			Trace.Println(err.Error())
			return
		}
		if bizData2, err = utf8ToGbk(bizData); err != nil {
			Trace.Println(err.Error())
			return
		}
		params.Set("biz_content", string(bizData2))
	}

	// 签名
	var sign string
	if sign, err = doSign(params, a.PrivateRSA); err != nil {
		Trace.Println(err.Error())
		return
	} else {
		params.Set("sign", sign)
	}

	var body []byte
	if body, err = a.request(params); err != nil {
		Trace.Println(err.Error())
		return
	}

	respStr := string(body)
	Trace.Println("response data======" + respStr)

	// 最初的数据..应该是 k v sign
	r := map[string]interface{}{}
	if err = json.Unmarshal(body, &r); err != nil {
		Trace.Println(err.Error())
		return
	}

	var data string
	//把 method 转换为 key
	respKey := strings.Replace(_api.getMethod(), ".", "_", -1) + "_response"
	// 有时候返回这个错误,替换为
	respStr = strings.Replace(respStr, "error_response", respKey, -1)
	//原始签名
	if v, pass := doVerifySign(respStr, fmt.Sprintf("%v", r["sign"]), respKey, a.PublicRSA); !pass {
		Trace.Println("verify failed")
	} else {
		data = v
	}

	var newBody []byte
	if newBody, err = gbkToUtf8([]byte(data)); err != nil {
		Trace.Println(err.Error())
		return
	}
	Trace.Println(string(newBody))

	if ptrResp == nil {
		ptrResp = map[string]interface{}{}
	}

	// 转义 GBK
	if err = json.Unmarshal(newBody, ptrResp); err != nil {
		Trace.Println(err.Error())
		return
	}

	return
}

/**
用于图片上传
*/
func (a *AlipayClient) ExecuteImageUpload(_api apiI, fileData []byte, fileName string, ptrResp interface{}) (err error) {
	// check
	if err = valid(a); err != nil {
		Trace.Println(err.Error())
		return
	}

	//系统参数
	params := a.initParams(_api.getMethod(), _api.getAppAuthToken())

	// 签名
	var sign string
	if sign, err = doSign(params, a.PrivateRSA); err != nil {
		Trace.Println(err.Error())
		return
	} else {
		params.Set("sign", sign)
	}

	var body []byte
	if body, err = a.requestMultiPart(params, fileData, fileName); err != nil {
		Trace.Println(err.Error())
		return
	}

	respStr := string(body)
	Trace.Println("response data======" + respStr)

	// 最初的数据..应该是 k v sign
	r := map[string]interface{}{}
	if err = json.Unmarshal(body, &r); err != nil {
		Trace.Println(err.Error())
		return
	}

	var data string
	//把 method 转换为 key
	respKey := strings.Replace(_api.getMethod(), ".", "_", -1) + "_response"
	// 有时候返回这个错误,替换为
	respStr = strings.Replace(respStr, "error_response", respKey, -1)
	//原始签名
	if v, pass := doVerifySign(respStr, fmt.Sprintf("%v", r["sign"]), respKey, a.PublicRSA); !pass {
		Trace.Println("verify failed")
	} else {
		data = v
	}

	var newBody []byte
	if newBody, err = gbkToUtf8([]byte(data)); err != nil {
		Trace.Println(err.Error())
		return
	}
	Trace.Println(string(newBody))

	if ptrResp == nil {
		ptrResp = map[string]interface{}{}
	}

	// 转义 GBK
	if err = json.Unmarshal(newBody, ptrResp); err != nil {
		Trace.Println(err.Error())
		return
	}

	return
}

func (a *AlipayClient) requestMultiPart(params url.Values, fileData []byte, fileName string) (data []byte, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	var formFile io.Writer
	if formFile, err = writer.CreateFormFile("image_content", fileName); err != nil {
		Trace.Println(err.Error())
		return
	} else if _, err = formFile.Write(fileData); err != nil {
		Trace.Println(err.Error())
		return
	}

	for key := range params {
		_ = writer.WriteField(key, params.Get(key))
	}

	if err = writer.Close(); err != nil {
		Trace.Println(err.Error())
		return
	}

	client := &http.Client{}
	client.Timeout = time.Duration(a.timeout) * time.Second

	var req *http.Request
	if req, err = http.NewRequest("POST", "https://openapi.alipay.com/gateway.do", body); err != nil {
		Trace.Println(err.Error())
		return
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		Trace.Println(err.Error())
		return
	}
	if resp.Body == nil {
		err = errors.New("response body is nil")
		Trace.Println(err.Error())
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	data, err = ioutil.ReadAll(resp.Body)
	return
}

func (a *AlipayClient) request(params url.Values) (data []byte, err error) {
	// 请求
	var res *http.Response
	if res, err = http.PostForm("https://openapi.alipay.com/gateway.do", params); err != nil {
		Trace.Println(err.Error())
		return
	}
	if res.Body == nil {
		err = errors.New("response body is null")
		Trace.Println(err.Error())
		return
	}

	defer func() {
		_ = res.Body.Close()
	}()

	data, err = ioutil.ReadAll(res.Body)
	return
}
