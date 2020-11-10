package open_alipay

/*
支付宝
公共请求参数
*/
type CommonParams struct {
	AppId        string `json:"app_id"`
	Method       string `json:"method"`
	Format       string `json:"format"`
	Charset      string `json:"charset"`
	SignType     string `json:"sign_type"`
	Sign         string `json:"sign"`
	TimeStamp    string `json:"timestamp"`
	Version      string `json:"version"`
	AppAuthToken string `json:"app_auth_token"`
}
