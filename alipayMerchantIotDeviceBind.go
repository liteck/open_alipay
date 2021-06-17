package open_alipay

/**
LOT 设备绑定接口
alipay.merchant.iot.device.bind
IoT设备绑定门店
*/
type AlipayMerchantIotDeviceBind struct {
	api
}

func (a AlipayMerchantIotDeviceBind) getMethod() string {
	return "alipay.merchant.iot.device.bind"
}

func (a AlipayMerchantIotDeviceBind) getAppAuthToken() string {
	return a.appAuthToken
}

func (a AlipayMerchantIotDeviceBind) getReq() interface{} {
	return a.input
}

type ReqAlipayMerchantIotDeviceBind struct {
	/*
		可选方式 [ID,SN]。
		ID-使用biztid作为设备唯一识别标识；
		SN-使用supplier_id、device_sn联合作为设备唯一识别标识。
		由于不同机型的supplier_id不同，推荐使用 ID 。
	*/
	DeviceIdType string `json:"device_id_type" validate:"required,eq=ID|eq=SN"`
	// 设备id 特殊可选 device_id_type填写了“ID”则可选
	BizTid string `json:"biz_tid,omitempty" validate:"omitempty"`
	// 设备供应商ID ，device_id_type 为 SN 时填写。需注意不同机型的供应商ID可能不同。
	SupplierId string `json:"supplier_id" validate:"required,eq=201901111100635561"`
	// 请填写设备的SN
	DeviceSn string `json:"device_sn" validate:"required"`
	// 商户类型，直连商户填写direct，间连商户填写indirect
	MerchantType string `json:"merchant_type" validate:"required,eq=direct|eq=indirect"`
	// 特殊可选.merchant_id_type为间连indirect时，商户smid已升级到M4等级，关联的pid。
	Pid string `json:"pid,omitempty" validate:"omitempty"`
	// 直连场景不填，间连场景填写商户收单smid
	Smid string `json:"smid" validate:"required"`
	// 店铺ID
	ShopId string `json:"shop_id" validate:"required"`
}

type RespAlipayMerchantIotDeviceBind struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
