package open_alipay

/**
LOT 设备绑定关系查询接口
alipay.commerce.iot.device.bind.query
设备绑定关系查询接口
*/
type AlipayCommerceIotDeviceBindQuery struct {
	api
}

func (a AlipayCommerceIotDeviceBindQuery) getMethod() string {
	return "alipay.commerce.iot.device.bind.query"
}

func (a AlipayCommerceIotDeviceBindQuery) getAppAuthToken() string {
	return a.appAuthToken
}

func (a AlipayCommerceIotDeviceBindQuery) getReq() interface{} {
	return a.input
}

type ReqAlipayCommerceIotDeviceBindQuery struct {
	//请填写MINI_APP
	AppType string `json:"app_type" validate:"required,eq=MINI_APP"`
	// RUYI_LITE
	MiniAppId string `json:"mini_app_id" validate:"required,eq=RUYI_LITE"`
	// SN
	DeviceIdType string `json:"device_id_type" validate:"required,eq=SN"`
	// 设备id 特殊可选 device_id_type填写了“SN”则可选
	BizTid string `json:"biz_tid,omitempty" validate:"omitempty"`
	// 特殊可选 device_id_type填写了“SN”则必填 设备供应商ID
	SupplierId string `json:"supplier_id" validate:"required,eq=201901111100635561"`
	// 请填写设备的SN
	DeviceSn string `json:"device_sn" validate:"required"`
}

type RespAlipayCommerceIotDeviceBindQuery struct {
	Code       string `json:"code"`
	Msg        string `json:"msg"`
	SubCode    string `json:"sub_code"`
	SubMsg     string `json:"sub_msg"`
	RetCode    string `json:"ret_code"`
	RetMessage string `json:"ret_message"`
}
