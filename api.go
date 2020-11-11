package open_alipay

import (
	"encoding/json"
)

type api interface {
	method() string
	params() (interface{}, interface{}, error)
	response() interface{}
}

type base struct {
}

func String(v interface{}) string {
	x, _ := json.Marshal(v)
	return string(x)
}


//func (a *AlipayApi) biz_to_string(b bizInterface) (string, error) {
//	if err := b.valid(); err != nil {
//		return "", err
//	}
//	content := ""
//	if v, err := json.Marshal(&b); err != nil {
//		return "", err
//	} else {
//		content = string(v)
//	}
//
//	// temp_map := map[string]interface{}{
//	// 	"biz": content,
//	// }
//
//	// if v, err := json.Marshal(&temp_map); err != nil {
//	// 	return "", err
//	// } else {
//	// 	content = string(v)
//	// }
//	// return content[8 : len(content)-2], nil
//	/**
//	20170928: 一直好好的.调试线上商品名字的时候.乱码.加上这个转码就好了.
//	通不通用有待验证
//	*/
//	content = tools.ConvertUTF2GBK(content)
//	return content, nil
//}
//

//func (o *OpenApi) verifySign(s, origin_sign, method_key string) bool {
//	if o.params.Method == "alipay.user.userinfo.share" {
//		//这里是要转义后校验的.NND
//		s = strings.Replace(s, "\\", "", -1)
//	}
//	sign_start_index := strings.Index(s, ",\"sign\"")
//	if sign_start_index == -1 {
//		return false
//	}
//	tobe_signed := s[4+len(method_key) : sign_start_index]
//	block, _ := pem.Decode(o.secret.AliPubRSA)
//	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		fmt.Printf("Failed to parse RSA public key: %s\n", err)
//		return false
//	}
//	rsaPub, _ := pub.(*rsa.PublicKey)
//	t := sha1.New()
//	io.WriteString(t, tobe_signed)
//	digest := t.Sum(nil)
//	data, err := base64.StdEncoding.DecodeString(origin_sign)
//	if err != nil {
//		fmt.Println("DecodeString sig error, reason: ", err)
//		return false
//	}
//	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
//	if err != nil {
//		fmt.Println("Verify sig error, reason: ", err)
//		return false
//	}
//
//	return true
//}
//
//func (o *OpenApi) request(m map[string]interface{}) (string, error) {
//	url_link := "https://openapi.alipay.com/gateway.do"
//	if conf.SandBoxEnable {
//		url_link = "https://openapi.alipaydev.com/gateway.do"
//	}
//	logs.Debug(fmt.Sprintf("==[request params]==[%s]", url_link))
//	http_request := httplib.Post(url_link)
//	tmp_string := ""
//	for k, _ := range m {
//		value := fmt.Sprintf("%v", m[k])
//		if value != "" {
//			http_request.Param(k, value)
//			tmp_string = tmp_string + k + "=" + value + "\t"
//		}
//	}
//	if len(o.filepath) > 0 {
//		//如果定义了这个文件..则上传
//		http_request.PostFile("image_content", o.filepath)
//		logs.Error(o.filepath)
//	}
//	logs.Debug(fmt.Sprintf("==[reuest params]==[%s]", tmp_string))
//	var string_result string
//	if v, err := http_request.String(); err != nil {
//		return "", err
//	} else {
//		string_result = v
//
//	}
//	return string_result, nil
//}
//
//func (o *OpenApi) Run(resp responseInterface) error {
//
//
//	if v, err := o.sign(tobe_sign); err != nil {
//		return err
//	} else if len(v) == 0 {
//		return ErrSign
//	} else {
//		__sign = v
//	}
//	logs.Debug(fmt.Sprintf("==[sign result]==[%s]", __sign))
//	m["sign"] = __sign
//	//准备请求
//	result_string := ""
//	if v, err := o.request(m); err != nil {
//		return err
//	} else {
//		result_string = v
//		//logs.Debug(fmt.Sprintf("==[response]==[gbk encode]:[%s]", result_string))
//	}
//
//	//把 method 转换为 key
//	method_key := strings.Replace(o.params.Method, ".", "_", -1)
//	method_key += "_response"
//
//	result_string = strings.Replace(result_string, "error_response", method_key, -1)
//
//	//解析结果
//	resp_map := map[string]interface{}{}
//	if err := json.Unmarshal([]byte(result_string), &resp_map); err != nil {
//		return err
//	}
//	//原始签名
//	if v, ok := resp_map["sign"].(string); ok && len(v) > 0 {
//		if pass := o.verifySign(result_string, v, method_key); !pass {
//			return ErrVerifySign
//		}
//	}
//
//	//转码
//	result_string = tools.ConvertGBK2UTF(result_string)
//	logs.Debug(fmt.Sprintf("==[response]==[utf-8 encode]:[%s]", result_string))
//	// 不知道以前为什么要加这个转义...先去掉
//	// result_string = strings.Replace(result_string, "\\", "", -1)
//	//logs.Debug(fmt.Sprintf("==[response]==[transferred]:[%s]", result_string))
//	if err := json.Unmarshal([]byte(result_string), &resp_map); err != nil {
//		return err
//	}
//	//把需要的内容再次换成string
//	if v, err := json.Marshal(resp_map[method_key]); err != nil {
//		return err
//	} else {
//		result_string = string(v)
//	}
//
//	if err := json.Unmarshal([]byte(result_string), resp); err != nil {
//		return err
//	}
//	return nil
//}
