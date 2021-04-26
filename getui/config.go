package getui

const (
	bashUrl = "https://restapi.getui.com/v2"
)

type ConfigParam struct {
	AppID        string       `json:"app_id,omitempty" form:"app_id"`
	AppKey       string       `json:"app_key,omitempty" form:"app_key"`
	MasterSecret string       `json:"master_secret,omitempty" form:"master_secret"`
	baseUrl      string       `json:"base_url,omitempty" form:"base_url"`
	AuthSignResp AuthSignResp `json:"auth_sign_resp,omitempty" form:"auth_sign_resp"`
}
