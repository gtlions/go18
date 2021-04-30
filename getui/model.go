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
	*Resp
	Data *AuthSignRespData `json:"data"`
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

type Settings struct {
	TTL int `json:"ttl"`
}
type Audience struct {
	Cid []string `json:"cid"`
}
type Notification struct {
	Title        string `json:"title,omitempty" form:"title"`
	Body         string `json:"body,omitempty" form:"body"`
	BigText      string `json:"big_text,omitempty" form:"big_text"`
	BigImage     string `json:"big_image,omitempty" form:"big_image"`
	Logo         string `json:"logo,omitempty" form:"logo"`
	LogoURL      string `json:"logo_url,omitempty" form:"logo_url"`
	ChannelID    string `json:"channel_id,omitempty" form:"channel_id"`
	ChannelName  string `json:"channel_name,omitempty" form:"channel_name"`
	ChannelLevel int    `json:"channel_level,omitempty" form:"channel_level"`
	ClickType    string `json:"click_type,omitempty" form:"click_type"`
	Intent       string `json:"intent,omitempty" form:"intent"`
	URL          string `json:"url,omitempty" form:"url"`
	Payload      string `json:"payload,omitempty" form:"payload"`
	NotifyID     int    `json:"notify_id,omitempty" form:"notify_id"`
	RingName     string `json:"ring_name,omitempty" form:"ring_name"`
	DadgeAddNum  int    `json:"badge_add_num,omitempty" form:"dadge_add_num"`
}

type Revoke struct {
	OldTaskID string `json:"old_task_id,omitempty" form:"old_task_id"`
	Force     bool   `json:"force,omitempty" form:"force"`
}

type TransmissionCst struct {
	Title   string `json:"title,omitempty" form:"title"`
	Content string `json:"content,omitempty" form:"content"`
	Payload string `json:"payload,omitempty" form:"payload"`
}

type Message struct {
	Duration     string        `json:"duration,omitempty"`
	Notification *Notification `json:"notification,omitempty"`
	Transmission string        `json:"transmission,omitempty"`
	Revoke       *Revoke       `json:"revoke,omitempty"`
}

type Alert struct {
	Title           string   `json:"title,omitempty" form:"title"`
	Body            string   `json:"body,omitempty" form:"body"`
	ActionLocKey    string   `json:"action-loc-key,omitempty" form:"action-loc-key"`
	LocKey          string   `json:"loc-key,omitempty" form:"loc-key"`
	LocArgs         []string `json:"loc-args,omitempty" form:"loc-args"`
	LaunchImage     string   `json:"launch-image,omitempty" form:"launch-image"`
	TitleLocKey     string   `json:"title-loc-key,omitempty" form:"title-loc-key"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty" form:"title-loc-args"`
	Subtitle        string   `json:"subtitle,omitempty" form:"subtitle"`
	SubtitleLocKey  string   `json:"subtitle-loc-key,omitempty" form:"subtitle-loc-key"`
	SubtitleLocArgs string   `json:"subtitle-loc-args,omitempty" form:"subtitle-loc-args"`
}

type Aps struct {
	Alert            *Alert `json:"alert,omitempty" form:"alert"`
	ContentAvailable int    `json:"content-available" form:"content-available"`
	Sound            string `json:"sound,omitempty" form:"sound"`
	Category         string `json:"category,omitempty" form:"category"`
	ThreadID         string `json:"thread-id,omitempty" form:"thread-id"`
}

type Multimedia struct {
	Url      string `json:"url,omitempty" form:"url"`
	Type     int    `json:"type,omitempty" form:"type"`
	OnlyWifi bool   `json:"only_wifi,omitempty" form:"only_wifi"`
}

type IOS struct {
	Type           string        `json:"type,omitempty" form:"type"`
	Aps            *Aps          `json:"aps,omitempty" form:"aps"`
	AutoBadge      string        `json:"auto_badge,omitempty" form:"auto_badge"`
	Payload        string        `json:"payload,omitempty" form:"payload"`
	Multimedia     []*Multimedia `json:"multimedia,omitempty" form:"multimedia"`
	ApnsCollapseID string        `json:"apns-collapse-id,omitempty" form:"apns-collapse-id"`
}
type Android struct {
}

type Channel struct {
	IOS     *IOS     `json:"ios,omitempty" form:"ios"`
	Android *Android `json:"android,omitempty" form:"android"`
}

type PushReq struct {
	RequestID   string    `json:"request_id"`
	Settings    *Settings `json:"settings"`
	Audience    *Audience `json:"audience"`
	PushMessage *Message  `json:"push_message"`
	PushChannel *Channel  `json:"push_channel"`
}
type BatchPush struct {
	IsAsync bool      `json:"is_async,omitempty" form:"is_async"`
	MsgList []PushReq `json:"msg_list,omitempty" form:"msg_list"`
}
type PushRespData struct {
	TaskID string `json:"task_id"`
	Status string `json:"status"`
}

type PushResp struct {
	*Resp
	*PushRespData
}

type GetAliasRespData struct {
	Alias string `json:"alias"`
}

type GetAliasResp struct {
	*Resp
	Data *GetAliasRespData `json:"data"`
}
