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

//var testToken = "202011BB36999f3d9ca440a6bb3fd265fc9d4X35"
var testToken = ""

func Test_alipay_open_sp_image_upload(t *testing.T) {
	resp := RespAlipayOpenSpImageUpload{}
	upload := AlipayOpenSpImageUpload{}
	if err := alipay.ExecuteImageUpload(upload,"door.png", &resp); err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println(resp)
	}
}

func Test_AlipayOpenSpBlueSeaActivityCreate(t *testing.T) {
	req := ReqAlipayOpenSpBlueSeaActivityCreate{}
	req.BizScene = "BLUE_SEA_FOOD_APPLY"
	req.MerchantLogon = "387737151@qq.com"
	req.BusinessLic = "123"
	resp := RespAlipayOpenSpBlueSeaActivityCreate{}

	create := AlipayOpenSpBlueSeaActivityCreate{}
	create.SetAuthToken(testToken)
	create.SetParams(req)
	if err := alipay.Execute(create, &resp); err != nil {
		log.Println(err.Error())
		return
	} else {
		log.Println(resp)
	}
}
