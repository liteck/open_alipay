package open_alipay

/**
图片上传接口
alipay.open.sp.image.upload
图片上传接口
*/
type AlipayOpenSpImageUpload struct {
	baseApi
}

func (a AlipayOpenSpImageUpload) getMethod() string {
	return "alipay.open.sp.image.upload"
}

func (a AlipayOpenSpImageUpload) getReq() (biz interface{}, form interface{}) {
	return nil, a.Input
}

type ReqAlipayOpenSpImageUpload struct {
	/**
	Byte_array
	图片二进制字节流，最大为10M
	二进制字节流
	*/
	ImageContent []byte `json:"image_content,omitempty"`
}

type RespAlipayOpenSpImageUpload struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
	//图片在文件存储平台的标识	64
	ImageId string `json:"image_id"`
}
