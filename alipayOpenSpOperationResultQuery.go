package open_alipay

/**
向商户发起代运营操作
alipay.open.sp.operation.apply
向商户发起代运营操作
*/
type AlipayOpenSpOperationResultQuery struct {
	api
}

func (a AlipayOpenSpOperationResultQuery) getMethod() string {
	return "alipay.open.sp.operation.result.query"
}

func (a AlipayOpenSpOperationResultQuery) getAppAuthToken() string {
	return a.appAuthToken
}

func (a AlipayOpenSpOperationResultQuery) getReq() interface{} {
	return a.input
}

type ReqAlipayOpenSpOperationResultQuery struct {
	//代运营操作类型。取值如下：
	//ACCOUNT_BIND：账号绑定，仅对于间连商户。
	//OPERATION_AUTH：代运营授权，支持间连及直连商户。
	OperateType string `json:"operate_type,omitempty" validate:"required,eq=ACCOUNT_BIND|eq=OPERATION_AUTH"`
	/**
	支付宝操作批次号
	*/
	BatchNo string `json:"batch_no"`
}

type RespAlipayOpenSpOperationResultQuery struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
	/**
	轮询结果。SUCCESS代表成功;PROCESS处理中
	*/
	HandleStatus string `json:"handle_status"`
	//被代运营的商户号。
	MerchantNo string `json:"merchant_no"`
	//被代运营者为间连商户时有值，此时返回绑定的支付宝账号。
	BindUserId string `json:"bind_user_id"`
}
