package getui

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// SetUserAlias 绑定别名
func SetUserAlias(config ConfigParam, userAliasReq SetUserAliasReq) (rsp Resp, err error) {
	client := http.DefaultClient
	payload, _ := json.Marshal(userAliasReq)
	req, _ := http.NewRequest("POST", config.baseUrl+"/user/alias", bytes.NewBuffer(payload))
	req.Header.Add("Token", config.AuthSignResp.Data.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Charset", "utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return
	}
	return rsp, err
}

// GetUserAlias 根据cid查询别名
func GetUserAlias(config ConfigParam, cid string) (rsp GetAliasResp, err error) {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", config.baseUrl+"/user/alias/cid"+"/"+cid, nil)
	req.Header.Add("Token", config.AuthSignResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return
	}
	return rsp, err
}

// GetUserCid 根据别名查询cid
func GetUserCid(config ConfigParam, alias string) (rsp GetAliasResp, err error) {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", config.baseUrl+"/user/cid/alias"+"/"+alias, nil)
	req.Header.Add("Token", config.AuthSignResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return
	}
	return rsp, err
}

// UnSetUserAliasBatch 批量解绑别名
func UnSetUserAliasBatch(config ConfigParam, alias SetUserAliasReq) (rsp Resp, err error) {
	payload, _ := json.Marshal(alias)
	client := http.DefaultClient
	req, _ := http.NewRequest("DELETE", config.baseUrl+"/user/alias", bytes.NewBuffer(payload))
	req.Header.Add("Token", config.AuthSignResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return
	}
	return rsp, err
}

// UnSetUserAlias 解绑所有别名
func UnSetUserAlias(config ConfigParam, alias string) (rsp Resp, err error) {
	client := http.DefaultClient
	req, _ := http.NewRequest("DELETE", config.baseUrl+"/user/alias"+"/"+alias, nil)
	req.Header.Add("Token", config.AuthSignResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return
	}
	return rsp, err
}

// GetUserStatus 查询用户状态
func GetUserStatus(config ConfigParam, cid string) (rsp GetAliasResp, err error) {
	client := http.DefaultClient
	req, _ := http.NewRequest("GET", config.baseUrl+"/user/status"+"/"+cid, nil)
	req.Header.Add("Token", config.AuthSignResp.Data.Token)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(rspBody, &rsp)
	if err != nil {
		return
	}
	return rsp, err
}
