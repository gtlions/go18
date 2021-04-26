package getui

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gtlions/gos10i"
)

// TestAuth
func TestAuth(t *testing.T) {
	var (
		appID        = "xE1r97JijZ6Aj9CVjVvsa9a7"
		appKey       = "3EdESIR2ndtY9IfBfiLMCZQ7"
		masterSecret = "p3suROTV6f17flDjpIln1I1"
	)
	config := ConfigParam{AppID: appID, AppKey: appKey, MasterSecret: masterSecret}
	if err := GetAuthToken(&config); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(config, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}

func TestSetUserAlias(t *testing.T) {
	appID := "E1r97JijZ6Aj9dsCVjVvsa97"
	config := ConfigParam{}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1d726ce256wewr1063dbcc4ffed8d62533d762bb1b5066757d588c1291d37d"
	users := []UserAlias{}
	users = append(users, UserAlias{Cid: "71ee1265b15e35ds0f6a92876f485bb6f", Alias: "ldm"})
	userAliasReq := SetUserAliasReq{DataList: users}
	if rsp, err := SetUserAlias(config, userAliasReq); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}
func TestUnSetUserAlias(t *testing.T) {
	appID := "E1r97JijZs6Aj9Cfs23VjVvsa97"
	config := ConfigParam{}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ace1d726ce256106df3dbcc4ffed8d62533d762bb1b5066757d588c1291d37d"
	alias := "ldm"
	if rsp, err := UnSetUserAlias(config, alias); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}

func TestGetUserAlias(t *testing.T) {
	appID := "E123r97JijfZ6Aj9CVfjVvsa97"
	config := ConfigParam{}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1d726ce256asdss1063dbcc4ffed8d62533d7f62bb1b5066757d588c1291d37d"
	cid := "71ee1265b15e35ad0f6a92821376f485bb6f"
	if rsp, err := GetUserAlias(config, cid); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}

func TestPushToSingleCid(t *testing.T) {
	appID := "E1r97J2ijZ6Aj9CserVjVvsa97"
	config := ConfigParam{}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1sd726ce256st1063dbcc4ffed8d62533xwd762bb1b5066757d588c1291d37d"
	push := PushReq{}
	push.RequestID = gos10i.X2String(time.Now().UnixNano())
	push.Settings.TTL = 3600 * 1000
	push.Audience.Cid = []string{"71ee1265b15sde35d0f6a92876f485b5wb6f"}
	push.PushMessage = PushMessage{Notification: PushMessageNotification{Title: "测试推送", Body: time.Now().String(), ClickType: "url", URL: "https//:xxx"}}
	if rsp, err := PushToSingleCid(config, push); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}
