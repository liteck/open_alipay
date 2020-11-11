package open_alipay

import (
	"log"
)

/**
新蓝海活动申请单详情查询
alipay.open.sp.blueseaactivity.query
新蓝海活动申请单详情查询
*/
type AlipayOpenSpBlueSeaActivityQuery struct {
	base
	Req  ReqAlipayOpenSpBlueSeaActivityQuery
	Resp *RespAlipayOpenSpBlueSeaActivityQuery
}

func (a AlipayOpenSpBlueSeaActivityQuery) method() string {
	return "alipay.open.sp.blueseaactivity.query"
}

func (a AlipayOpenSpBlueSeaActivityQuery) params() (biz interface{}, other interface{}, err error) {
	if err = valid(&a.Req); err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}

	return a.Req, nil, nil
}

func (a AlipayOpenSpBlueSeaActivityQuery) response() interface{} {
	return a.Resp
}

type ReqAlipayOpenSpBlueSeaActivityQuery struct {
	//申请单Id
	OrderId string `json:"order_id,omitempty" validate:"required"`
}

type RespAlipayOpenSpBlueSeaActivityQuery struct {
	Sign     string `json:"sign"`
	Response struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
		/**
		蓝海活动的场景，
		包括直连餐饮（BLUE_SEA_FOOD_APPLY）、
		直连快消（BLUE_SEA_FMCG_APPLY）、
		间连餐饮（BLUE_SEA_FOOD_INDIRECT_APPLY）、
		间连快消（BLUE_SEA_FMCG_INDIRECT_APPLY）
		场景
		*/
		BizScene string `json:"biz_scene"`
		/*
		   参与蓝海活动的商户支付宝账号，
		   只有当参与直连蓝海活动场景（BLUE_SEA_FOOD_APPLY/BLUE_SEA_FMCG_APPLY）时必填，
		   间连场景可空
		*/
		MerchantLogon string `json:"merchant_logon"`
		/**
		参与蓝海活动的间连商户账号，
		只有当参与间连蓝海活动场景（BLUE_SEA_FOOD_INDIRECT_APPLY/BLUE_SEA_FMCG_INDIRECT_APPLY）时必填，
		直连场景可空
		*/
		SubMerchantId string `json:"sub_merchant_id"`
		/**
		申请单状态，状态机参考
		AUDITING:审核中，
		FAIL:报名失败，
		PASS:报名成
		*/
		Status string `json:"status"`
		/**
		营业执照，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		BusinessLic string `json:"business_lic"`
		/**
		餐饮服务许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		FoodServiceLic string `json:"food_service_lic"`
		/**
		食品卫生许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		FoodHealthLic string `json:"food_health_lic"`
		/**
		食品经营许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		FoodBusinessLic string `json:"food_business_lic"`
		/**
		食品流通许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		FoodCirculateLic string `json:"food_circulate_lic"`
		/**
		食品生产许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		FoodProductionLic string `json:"food_production_lic"`
		/**
		烟草专卖零售许可证，要求证件文本信息清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		TobaccoLic string `json:"tobacco_lic"`
		/**
		门头照，要求店铺外观照片清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		ShopEntrancePic string `json:"shop_entrance_pic"`
		/**
		店内照，要求店内照片清晰可见。 请上传照片fileId（传参明细参见应用场景说明）
		*/
		IndoorPic string `json:"indoor_pic"`
		/**
		省份编码
		*/
		ProvinceCode string `json:"province_code"`
		/**
		城市编码
		*/
		CityCode string `json:"city_code"`
		/**
		区县编码
		*/
		DistrictCode string `json:"district_code"`
		/**
		详细地址 万塘路18号黄龙时代广场B座
		*/
		Address string `json:"address"`
	} `json:"alipay_open_sp_blueseaactivity_query_response"`
}
