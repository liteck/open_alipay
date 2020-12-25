package open_alipay

import (
	"log"
	"testing"
)

var PublicRSA = []byte(`-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----`)

var PrivateRSA = []byte(`-----BEGIN RSA PRIVATE KEY-----
-----END RSA PRIVATE KEY-----`)

var alipay = AlipayClient{
	AppId:      "",
	UserId:     "",
	PublicRSA:  PublicRSA,
	PrivateRSA: PrivateRSA,
}

var testToken = ""

func Test_alipay_open_sp_image_upload(t *testing.T) {
	resp := RespAlipayOpenSpImageUpload{}
	upload := AlipayOpenSpImageUpload{}
	if err := alipay.ExecuteImageUpload(upload, "door.png", &resp); err != nil {
		Trace.Println(err.Error())
		return
	} else {
		Trace.Println(resp)
	}
}

func Test_AlipayOpenSpBlueSeaActivityCreate(t *testing.T) {
	req := ReqAlipayOpenSpBlueSeaActivityCreate{}
	req.BizScene = "BLUE_SEA_FOOD_APPLY"
	req.MerchantLogon = "387737151@qq.com"
	req.BusinessLic = "A*kgAFSpNEPk4AAAAAAAAAAAAADsF1AQ"
	req.ShopEntrancePic = "A*kgAFSpNEPk4AAAAAAAAAAAAADsF1AQ"
	req.IndoorPic = "A*kgAFSpNEPk4AAAAAAAAAAAAADsF1AQ"
	req.ProvinceCode = "310000"
	req.CityCode = "310100"
	req.DistrictCode = "310113"
	req.Address = "陆翔路1018弄6号101室-1"
	resp := RespAlipayOpenSpBlueSeaActivityCreate{}

	create := AlipayOpenSpBlueSeaActivityCreate{}
	create.SetAuthToken(testToken)
	create.SetParams(req)
	if err := alipay.Execute(create, &resp); err != nil {
		Trace.Println(err.Error())
		return
	}
}
