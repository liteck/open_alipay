package open_alipay

/**
向商户发起代运营操作
ant.merchant.expand.shop.page.query
向商户发起代运营操作
*/
type AntMerchantExpandShopPageQuery struct {
	api
}

func (a AntMerchantExpandShopPageQuery) getMethod() string {
	return "ant.merchant.expand.shop.page.query"
}

func (a AntMerchantExpandShopPageQuery) getAppAuthToken() string {
	return a.appAuthToken
}

func (a AntMerchantExpandShopPageQuery) getReq() interface{} {
	return a.input
}

type ReqAntMerchantExpandShopPageQuery struct {
	/**
	商户角色id，表示将要开的店属于哪个商户角色。
	对于直连开店场景，填写商户pid；
	对于间连开店场景（线上、线下、直付通），填写商户smid。
	特别说明：IoT设备三绑定场景统一填写商户pid
	*/
	IpRoleId string `json:"ip_role_id,omitempty" validate:"required"`
	//查询页数 1
	PageNum int `json:"page_num,omitempty" validate:"required"`
	//每页查询大小，限制100以内
	PageSize int `json:"page_size,omitempty" validate:"required"`
}

type RespAntMerchantExpandShopPageQuery struct {
	Code       string               `json:"code"`
	Msg        string               `json:"msg"`
	SubCode    string               `json:"sub_code"`
	SubMsg     string               `json:"sub_msg"`
	TotalPages int                  `json:"total_pages"`
	ShopInfos  []ShopQueryOpenApiVO `json:"shop_infos"`
}

type ShopQueryOpenApiVO struct {
	ShopId          string      `json:"shop_id"`
	BusinessAddress AddressInfo `json:"business_address"`
	ShopCategory    string      `json:"shop_category"`
	StoreId         string      `json:"store_id"`
	ShopType        string      `json:"shop_type"`
	ShopName        string      `json:"shop_name"`
	ContactPhone    string      `json:"contact_phone"`
	ContactMobile   string      `json:"contact_mobile"`
	ShopStatus      string      `json:"shop_status"`
}

type AddressInfo struct {
	CityCode     string `json:"city_code"`
	DistrictCode string `json:"district_code"`
	Address      string `json:"address"`
	ProvinceCode string `json:"province_code"`
	Poiid        string `json:"poiid"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	Type         string `json:"type"`
}
