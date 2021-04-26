package getui

type Resp struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
type AuthSignParam struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Appkey    string `json:"appkey"`
}

type AuthSignResp struct {
	Resp
	Data AuthSignRespData `json:"data"`
}
type AuthSignRespData struct {
	Expiretime string `json:"expire_time"`
	Token      string `json:"token"`
}

type SetUserAliasReq struct {
	DataList []UserAlias `json:"data_list"`
}
type UserAlias struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}

type PushSettings struct {
	TTL int `json:"ttl"`
}
type PushAudience struct {
	Cid []string `json:"cid"`
}
type PushMessageNotification struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	ClickType string `json:"click_type"`
	URL       string `json:"url"`
}
type PushMessage struct {
	Notification PushMessageNotification `json:"notification"`
}
type PushReq struct {
	RequestID   string       `json:"request_id"`
	Settings    PushSettings `json:"settings"`
	Audience    PushAudience `json:"audience"`
	PushMessage PushMessage  `json:"push_message"`
}

type PushRespData struct {
	TaskID string `json:"task_id"`
	Status string `json:"status"`
}

type PushResp struct {
	Resp
	PushRespData
}

type GetAliasRespData struct {
	Alias string `json:"alias"`
}

type GetAliasResp struct {
	Resp
	Data GetAliasRespData `json:"data"`
}
