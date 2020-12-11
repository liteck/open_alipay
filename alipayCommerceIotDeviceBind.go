package open_alipay

/**
LOT 设备绑定接口
alipay.commerce.iot.device.bind
设备绑定接口
*/
type AlipayCommerceIotDeviceBind struct {
	api
}

func (a AlipayCommerceIotDeviceBind) getMethod() string {
	return "alipay.commerce.iot.device.bind"
}

func (a AlipayCommerceIotDeviceBind) getAppAuthToken() string {
	return a.appAuthToken
}

func (a AlipayCommerceIotDeviceBind) getReq() interface{} {
	return a.input
}

type ReqAlipayCommerceIotDeviceBind struct {
	//请填写MINI_APP
	AppType string `json:"app_type" validate:"required,eq=MINI_APP"`
	// RUYI_LITE
	MiniAppId string `json:"mini_app_id" validate:"required,eq=RUYI_LITE"`
	// 应用ID，目前暂不使用，不要填
	ApplicationId string `json:"application_id,omitempty" validate:"omitempty"`
	// SN
	DeviceIdType string `json:"device_id_type" validate:"required,eq=SN"`
	// 设备id 特殊可选 device_id_type填写了“SN”则可选
	BizTid string `json:"biz_tid,omitempty" validate:"omitempty"`
	// 特殊可选 device_id_type填写了“SN”则必填 设备供应商ID
	SupplierId string `json:"supplier_id" validate:"required,eq=201901111100635561"`
	// 请填写设备的SN
	DeviceSn string `json:"device_sn" validate:"required"`
	// 受理商户的ISV在支付宝的pid ..pid..
	Source string `json:"source" validate:"required"`
	// 商户编号，由ISV定义，需要保证在ISV下唯一
	ExternalId string `json:"external_id" validate:"required"`
	// 区分商户ID类型，直连商户填写direct，间连商户填写indirect
	MerchantIdType string `json:"merchant_id_type" validate:"required,eq=direct|eq=indirect"`
	// 商户角色id。对于直连开店场景，填写商户pid；对于间连开店场景，填写商户smid。
	MerchantId string `json:"merchant_id" validate:"required"`
	// 店铺ID
	ShopId string `json:"shop_id" validate:"required"`
	// 外部门店id
	ExternalShopId string `json:"external_shop_id,omitempty" validate:"omitempty"`
	// 特殊可选.merchant_id_type为间连indirect时，商户smid已升级到M4等级，关联的pid。
	Pid string `json:"pid,omitempty" validate:"omitempty"`
	// isv应用id，isv的标识；用于标识ISV身份
	SpiAppId string `json:"spi_app_id" validate:"required"`
	// isv内部唯一标识商户、商户门店等字段信息； 可作为额外ISV自主对接的扩展字段使用
	TerminalBindInfo string `json:"terminal_bind_info" validate:"required"`
}

type RespAlipayCommerceIotDeviceBind struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"sub_code"`
	SubMsg     string `json:"sub_msg"`
	RetCode    string `json:"ret_code"`
	RetMessage string `json:"ret_message"`
}
