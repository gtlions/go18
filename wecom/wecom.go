package wecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gtlions/go18"
)

type msgContent struct {
	Content string `json:"content"`
}
type msgPlayload struct {
	Touser   string     `json:"touser,omitempty"`
	Toparty  string     `json:"toparty,omitempty"`
	Msgtype  string     `json:"msgtype"`
	Agentid  uint       `json:"agentid"`
	Text     msgContent `json:"text"`
	Markdown msgContent `json:"markdown"`
	Safe     int        `json:"safe"`
}

type RespAccessToken struct {
	Errcode     int    `json:"errcode" yaml:"errcode" form:"errcode" gorm:"errcode"`
	Errmsg      string `json:"errmsg" yaml:"errmsg" form:"errmsg" gorm:"errmsg"`
	AccessToken string `json:"access_token" yaml:"access_token" form:"access_token" gorm:"access_token"`
	ExpiresIn   int    `json:"expires_in" yaml:"expires_in" form:"expires_in" gorm:"expires_in"`
}
type RespSendMessage struct {
	Errcode        int    `json:"errcode"`
	Errmsg         string `json:"errmsg"`
	Invaliduser    string `json:"invaliduser,omitempty" yaml:"invaliduser" form:"invaliduser" gorm:"invaliduser"`
	Invalidparty   string `json:"invalidparty,omitempty" yaml:"invalidparty" form:"invalidparty" gorm:"invalidparty"`
	Invalidtag     string `json:"invalidtag,omitempty" yaml:"invalidtag" form:"invalidtag" gorm:"invalidtag"`
	Unlicenseduser string `json:"unlicenseduser,omitempty" yaml:"unlicenseduser" form:"unlicenseduser" gorm:"unlicenseduser"`
	Msgid          string `json:"msgid,omitempty" yaml:"msgid" form:"msgid" gorm:"msgid"`
	ResponseCode   string `json:"response_code,omitempty" yaml:"response_code" form:"response_code" gorm:"response_code"`
}

type Wecom struct {
	CorpID     string `config:"corpid"`
	AgentID    uint   `config:"agentid"`
	CorpSecret string `config:"corpsecret"`
	ToUser     string `config:"touser"`
	ToParty    string `config:"toparty"`
	// Message         []string `config:"message"`
	RespAccessToken *RespAccessToken
	RespSendMessage *RespSendMessage
}

func (e *Wecom) GetAccessToken() error {
	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + e.CorpID + "&corpsecret=" + e.CorpSecret
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		e.RespAccessToken.Errcode = -1
		e.RespAccessToken.Errmsg = fmt.Sprintf("newrequest wechat access token failed with: %v", err)
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		e.RespAccessToken.Errcode = -1
		e.RespAccessToken.Errmsg = fmt.Sprintf("request wechat access token failed with: %v", err)
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e.RespAccessToken.Errcode = -1
		e.RespAccessToken.Errmsg = fmt.Sprintf("read wechat access token failed with: %v", err)
		return err
	}
	err = json.Unmarshal(body, &e.RespAccessToken)
	if err != nil {
		e.RespAccessToken.Errcode = -1
		e.RespAccessToken.Errmsg = fmt.Sprintf("decode wechat access token failed with: %v", err)
		return err
	}
	if e.RespAccessToken.Errcode != 0 {
		e.RespAccessToken.Errmsg = fmt.Sprintf("get wechat access token failed with: %v", e.RespAccessToken.Errmsg)
		return fmt.Errorf(e.RespAccessToken.Errmsg)
	}
	e.RespAccessToken.ExpiresIn = go18.XRandomIntRange(1, 9999)
	return nil
}

func (e *Wecom) SendText(msg []string) error {
	client := &http.Client{}
	for _, v := range msg {
		// msgSend := make_message_text(e.ToUser, e.ToParty, e.AgentID, v)
		msgSend := make_message_md(e.ToUser, e.ToParty, e.AgentID, v)
		url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + e.RespAccessToken.AccessToken
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(msgSend)))
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("newrequest wechat send failed with: %v", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("request wechat send failed with: %v", err)
			return err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("read wechat send failed with: %v", err)
			return err
		}
		err = json.Unmarshal(body, &e.RespSendMessage)
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("decode wechat send result failed with: %v", err)
			return err
		}
		if e.RespSendMessage.Errcode != 0 {
			e.RespSendMessage.Errmsg = fmt.Sprintf("send wechat message failed with: %v", e.RespSendMessage.Errmsg)
			return fmt.Errorf(e.RespSendMessage.Errmsg)
		}
	}
	return nil
}

func (e *Wecom) SendMD(msg []string) error {
	client := &http.Client{}
	for _, v := range msg {
		// msgSend := make_message_text(e.ToUser, e.ToParty, e.AgentID, v)
		msgSend := make_message_md(e.ToUser, e.ToParty, e.AgentID, v)
		url := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + e.RespAccessToken.AccessToken
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(msgSend)))
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("newrequest wechat send failed with: %v", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("request wechat send failed with: %v", err)
			return err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("read wechat send failed with: %v", err)
			return err
		}
		err = json.Unmarshal(body, &e.RespSendMessage)
		if err != nil {
			e.RespSendMessage.Errcode = -1
			e.RespSendMessage.Errmsg = fmt.Sprintf("decode wechat send result failed with: %v", err)
			return err
		}
		if e.RespSendMessage.Errcode != 0 {
			e.RespSendMessage.Errmsg = fmt.Sprintf("send wechat message failed with: %v", e.RespSendMessage.Errmsg)
			return fmt.Errorf(e.RespSendMessage.Errmsg)
		}
	}
	return nil
}

func make_message_text(touser string, toparty string, agentid uint, content string) string {
	msg := msgPlayload{
		Touser:  touser,
		Toparty: toparty,
		Msgtype: "text",
		Agentid: agentid,
		Safe:    0,
		Text: struct {
			Content string `json:"content"`
		}{Content: content},
	}
	sed_msg, _ := json.Marshal(msg)
	return string(sed_msg)
}

func make_message_md(touser string, toparty string, agentid uint, content string) string {
	msg := msgPlayload{
		Touser:  touser,
		Toparty: toparty,
		Msgtype: "markdown",
		Agentid: agentid,
		Markdown: struct {
			Content string `json:"content"`
		}{Content: content},
	}
	sed_msg, _ := json.Marshal(msg)
	return string(sed_msg)
}
