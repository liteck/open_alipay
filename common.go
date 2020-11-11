package open_alipay

//公共请求参数
type CommonParams struct {
	AppId    string `json:"app_id" validate:"required"`
	Method   string `json:"method" validate:"required"`
	Format   string `json:"format" validate:"required,eq=JSON"`
	Charset  string `json:"charset" validate:"required,eq=utf-8|eq=gbk|eq=gb2312"`
	SignType string `json:"sign_type" validate:"required,eq=RSA2|eq=RSA"`
	//Sign         string `json:"sign" validate:"required"`
	TimeStamp    string `json:"timestamp" validate:"required"`
	Version      string `json:"version" validate:"required,eq=1.0"`
	AppAuthToken string `json:"app_auth_token,omitempty" validate:"omitempty"`
}

func (c *CommonParams) Valid() (err error) {
	return valid(c)
}
