package open_alipay

import (
	"errors"
	"log"
)

/**
新蓝海活动报名申请
alipay.open.sp.blueseaactivity.create
新蓝海活动报名申请
*/
type AlipayOpenSpBlueSeaActivityCreate struct {
	base
	Req  ReqAlipayOpenSpBlueSeaActivityCreate
	Resp *RespAlipayOpenSpBlueSeaActivityCreate
}

func (a AlipayOpenSpBlueSeaActivityCreate) method() string {
	return "alipay.open.sp.blueseaactivity.create"
}

func (a AlipayOpenSpBlueSeaActivityCreate) params() (biz interface{}, other interface{}, err error) {
	if err = valid(&a.Req); err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	if (a.Req.BizScene == "BLUE_SEA_FOOD_APPLY" || a.Req.BizScene == "BLUE_SEA_FMCG_APPLY") && a.Req.MerchantLogon == "" {
		err = errors.New("参与直连蓝海活动场景时,商户支付宝账号必填")
		return
	}

	if (a.Req.BizScene == "BLUE_SEA_FOOD_INDIRECT_APPLY" || a.Req.BizScene == "BLUE_SEA_FMCG_INDIRECT_APPLY") && a.Req.SubMerchantId == "" {
		err = errors.New("参与间连蓝海活动场景时,间联商户号必填")
		return
	}

	return a.Req, nil, nil
}

func (a AlipayOpenSpBlueSeaActivityCreate) response() interface{} {
	return a.Resp
}

type ReqAlipayOpenSpBlueSeaActivityCreate struct {
	/**
	蓝海活动的场景，
	包括直连餐饮（BLUE_SEA_FOOD_APPLY）、
	直连快消（BLUE_SEA_FMCG_APPLY）、
	间连餐饮（BLUE_SEA_FOOD_INDIRECT_APPLY）、
	间连快消（BLUE_SEA_FMCG_INDIRECT_APPLY）
	场景
	*/
	BizScene string `json:"biz_scene" validate:"required,eq=BLUE_SEA_FOOD_APPLY|eq=BLUE_SEA_FMCG_APPLY|eq=BLUE_SEA_FOOD_INDIRECT_APPLY|eq=BLUE_SEA_FMCG_INDIRECT_APPLY"`
	/*
	   参与蓝海活动的商户支付宝账号，
	   只有当参与直连蓝海活动场景（BLUE_SEA_FOOD_APPLY/BLUE_SEA_FMCG_APPLY）时必填，
	   间连场景可空
	*/
	MerchantLogon string `json:"merchant_logon" validate:"omitempty"`
	/**
	参与蓝海活动的间连商户账号，
	只有当参与间连蓝海活动场景（BLUE_SEA_FOOD_INDIRECT_APPLY/BLUE_SEA_FMCG_INDIRECT_APPLY）时必填，
	直连场景可空
	*/
	SubMerchantId string `json:"sub_merchant_id" validate:"omitempty"`
	/**
	营业执照，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	BusinessLic string `json:"business_lic" validate:"omitempty"`
	/**
	餐饮服务许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodServiceLic string `json:"food_service_lic" validate:"omitempty"`
	/**
	食品卫生许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodHealthLic string `json:"food_health_lic" validate:"omitempty"`
	/**
	食品经营许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodBusinessLic string `json:"food_business_lic" validate:"omitempty"`
	/**
	食品流通许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodCirculateLic string `json:"food_circulate_lic" validate:"omitempty"`
	/**
	食品生产许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	FoodProductionLic string `json:"food_production_lic" validate:"omitempty"`
	/**
	烟草专卖零售许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	TobaccoLic string `json:"tobacco_lic" validate:"omitempty"`
	/**
	门头照，要求店铺外观照片清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	ShopEntrancePic string `json:"shop_entrance_pic" validate:"omitempty"`
	/**
	店内照，要求店内照片清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
	*/
	IndoorPic string `json:"indoor_pic" validate:"omitempty"`
	/**
	省份编码
	*/
	ProvinceCode string `json:"province_code" validate:"omitempty"`
	/**
	城市编码
	*/
	CityCode string `json:"city_code" validate:"omitempty"`
	/**
	区县编码
	*/
	DistrictCode string `json:"district_code" validate:"omitempty"`
	/**
	详细地址 万塘路18号黄龙时代广场B座
	*/
	Address string `json:"address" validate:"omitempty"`
}

type RespAlipayOpenSpBlueSeaActivityCreate struct {
	Sign     string `json:"sign"`
	Response struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
		//申请单Id
		OrderId string `json:"order_id"`
	} `json:"alipay_open_sp_blueseaactivity_create_response"`
}
