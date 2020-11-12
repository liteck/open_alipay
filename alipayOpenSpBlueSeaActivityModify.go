package open_alipay

/**
新蓝海活动报名申请单修改
alipay.open.sp.blueseaactivity.modify
新蓝海活动报名申请单修改
*/
type AlipayOpenSpBlueSeaActivityModify struct {
	baseApi
}

func (a AlipayOpenSpBlueSeaActivityModify) getMethod() string {
	return "alipay.open.sp.blueseaactivity.modify"
}

func (a AlipayOpenSpBlueSeaActivityModify) getReq() (biz interface{}, form interface{}) {
	return a.Input, nil
}

type ReqAlipayOpenSpBlueSeaActivityModify struct {
	//申请单Id
	OrderId string `json:"order_id,omitempty" validate:"required"`
	/**
	营业执照，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	BusinessLic string `json:"business_lic,omitempty" validate:"omitempty"`
	/**
	餐饮服务许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodServiceLic string `json:"food_service_lic,omitempty" validate:"omitempty"`
	/**
	食品卫生许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodHealthLic string `json:"food_health_lic,omitempty" validate:"omitempty"`
	/**
	食品经营许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodBusinessLic string `json:"food_business_lic,omitempty" validate:"omitempty"`
	/**
	食品流通许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodCirculateLic string `json:"food_circulate_lic,omitempty" validate:"omitempty"`
	/**
	食品生产许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodProductionLic string `json:"food_production_lic,omitempty" validate:"omitempty"`
	/**
	烟草专卖零售许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	TobaccoLic string `json:"tobacco_lic,omitempty" validate:"omitempty"`
	/**
	门头照，要求店铺外观照片清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	ShopEntrancePic string `json:"shop_entrance_pic,omitempty" validate:"omitempty"`
	/**
	店内照，要求店内照片清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	IndoorPic string `json:"indoor_pic,omitempty" validate:"omitempty"`
	/**
	省份编码
	*/
	ProvinceCode string `json:"province_code,omitempty" validate:"omitempty"`
	/**
	城市编码
	*/
	CityCode string `json:"city_code,omitempty" validate:"omitempty"`
	/**
	区县编码
	*/
	DistrictCode string `json:"district_code,omitempty" validate:"omitempty"`
	/**
	详细地址 万塘路18号黄龙时代广场B座
	*/
	Address string `json:"address,omitempty" validate:"omitempty"`
}

type RespAlipayOpenSpBlueSeaActivityModify struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}
