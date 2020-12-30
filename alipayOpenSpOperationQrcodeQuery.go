package open_alipay

/**
向商户发起代运营操作
alipay.open.sp.operation.apply
向商户发起代运营操作
*/
type AlipayOpenSpOperationQrcodeQuery struct {
	api
}

func (a AlipayOpenSpOperationQrcodeQuery) getMethod() string {
	return "alipay.open.sp.operation.qrcode.query"
}

func (a AlipayOpenSpOperationQrcodeQuery) getAppAuthToken() string {
	return a.appAuthToken
}

func (a AlipayOpenSpOperationQrcodeQuery) getReq() interface{} {
	return a.input
}

type ReqAlipayOpenSpOperationQrcodeQuery struct {
	//外部操作流水，ISV自定义。每次操作需要确保唯一。
	OutBizNo string `json:"out_biz_no,omitempty" validate:"required"`
	//代运营操作类型。取值如下：
	//ACCOUNT_BIND：账号绑定，仅对于间连商户。
	//OPERATION_AUTH：代运营授权，支持间连及直连商户。
	OperateType string `json:"operate_type,omitempty" validate:"required,eq=ACCOUNT_BIND|eq=OPERATION_AUTH"`
	//接入的产品编号。 枚举如下：
	//操作类型为账号绑定时，填OPENAPI_BIND_DEFAULT。
	//操作类型为代运营授权时，填OPENAPI_AUTH_DEFAULT。
	AccessProductCode string `json:"access_product_code,omitempty" validate:"required,eq=OPENAPI_BIND_DEFAULT|eq=OPENAPI_AUTH_DEFAULT"`
	//支付宝商户号。间连、直连商户均支持，特别注意仅支持2088开头的间连商户。
	//若被代运营者是间连商户，则merchant_no必填。
	//若为直连商户，则merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。
	MerchantNo string `json:"merchant_no,omitempty" validate:"omitempty"`
	//支付宝登录账号。通常为手机号或者邮箱。
	//若被代运营者是间连商户，则alipay_account必填。
	//若为直连商户，则merchant_no和alipay_account不能同时为空，都有值优先取merchant_no。
	AlipayAccount string `json:"alipay_account,omitempty" validate:"omitempty"`
}

type RespAlipayOpenSpOperationQrcodeQuery struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
	/**
	支付宝操作批次号
	*/
	BatchNo string `json:"batch_no"`
	//二维码图片地址。urlEncode处理过
	QrCodeUrl string `json:"qr_code_url"`
}
