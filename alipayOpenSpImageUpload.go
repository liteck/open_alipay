package open_alipay

import "log"

/**
图片上传接口
alipay.open.sp.image.upload
图片上传接口
*/
type AlipayOpenSpImageUpload struct {
	base
	Req  ReqAlipayOpenSpImageUpload
	Resp *RespAlipayOpenSpImageUpload
}

func (a AlipayOpenSpImageUpload) method() string {
	return "alipay.open.sp.image.upload"
}

func (a AlipayOpenSpImageUpload) params() (biz interface{}, other interface{}, err error) {
	if err = valid(&a.Req); err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	return a.Req, nil, nil
}

func (a AlipayOpenSpImageUpload) response() interface{} {
	return a.Resp
}

type ReqAlipayOpenSpImageUpload struct {
	/**
	Byte_array
	图片二进制字节流，最大为10M
	二进制字节流
	*/
	ImageContent []byte `json:"image_content;omitempty"`
}

type RespAlipayOpenSpImageUpload struct {
	Sign     string `json:"sign"`
	Response struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_open_sp_image_upload_response"`
	//图片在文件存储平台的标识	64
	ImageId string `json:"image_id"`
}
