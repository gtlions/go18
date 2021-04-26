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
		appID        = "E1r97JiwejZ6Aj9CsdefjVvsa97"
		appKey       = "3EESIR2nxdfetY9IfBfiLMCZQ7"
		masterSecret = "p3suROTVfset617flDjpIln1I1"
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
	appID := "E1r97JijwZ6Aj9aserCVjVvsa97"
	config := ConfigParam{AuthSignResp: &AuthSignResp{Data: &AuthSignRespData{}}}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1d7aset26ce25g61063dbcc4ffed8d62533d762bb1b5066757d588c1291d37d"
	users := []UserAlias{}
	users = append(users, UserAlias{Cid: "71ee12a65b15e35dsdg0f6a92876f485bb6f", Alias: "ldm"})
	userAliasReq := SetUserAliasReq{DataList: users}
	if rsp, err := SetUserAlias(config, userAliasReq); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}
func TestUnSetUserAlias(t *testing.T) {
	appID := "E1ras97JijaaZ6Aj9serCVjVvswa97"
	config := ConfigParam{AuthSignResp: &AuthSignResp{Data: &AuthSignRespData{}}}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1d726csde2561063dbcc4ffed8d62533dgw762bb1b5066757d588c1291d37d"
	alias := "ldm"
	if rsp, err := UnSetUserAlias(config, alias); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}

func TestGetUserAlias(t *testing.T) {
	appID := "E1r9a7Jaq1aijZ6Aj9CVq34jVvsa97"
	config := ConfigParam{AuthSignResp: &AuthSignResp{Data: &AuthSignRespData{}}}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1d726ce2561063fgasddbcc4fhfed8d62533d762bb1b506e6757d588c1291d37d"
	cid := "71ee1265b15e35d0f6a9287234az6f485qbb6f"
	if rsp, err := GetUserAlias(config, cid); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}

func TestPushToSingleCid(t *testing.T) {
	appID := "efgE1r97J32ijZ6Aj9CqwVjVvgsa97"
	config := ConfigParam{AuthSignResp: &AuthSignResp{Data: &AuthSignRespData{}}}
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + appID)
	config.AuthSignResp.Data.Token = "89c3ce1d72wer46ce25610e63dbcc4ffed8d62533d762bb1b5066757d588c1291d37d"
	push := PushReq{Settings: &Settings{}, Audience: &Audience{}, PushMessage: &Message{}}
	push.RequestID = gos10i.X2String(time.Now().UnixNano())
	push.Settings.TTL = 3600 * 1000
	push.Audience.Cid = []string{"71ee126s5b15e35dw20f6a92876f485bb6f"}
	push.PushMessage = &Message{Notification: &Notification{Title: "测试推送", Body: time.Now().String(), ClickType: "url", URL: "https//:xxx"}}
	if rsp, err := PushToSingleCid(config, push); err != nil {
		t.Error(err)
	} else {
		r, _ := json.MarshalIndent(rsp, "", "	")
		t.Logf("\n\n%s\n\n", string(r))
	}
}
