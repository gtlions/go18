package getui

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetAuthToken(config *ConfigParam) (err error) {
	config.baseUrl = fmt.Sprintf(bashUrl + "/" + config.AppID)
	timestamp := time.Now().UnixNano() / 1e6
	client := http.DefaultClient
	parm := AuthSignParam{Sign: fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%d%s", config.AppKey, timestamp, config.MasterSecret)))), Timestamp: fmt.Sprintf("%d", timestamp), Appkey: config.AppKey}
	payload, _ := json.Marshal(parm)
	req, _ := http.NewRequest("POST", config.baseUrl+"/auth", bytes.NewBuffer(payload))
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
	err = json.Unmarshal(rspBody, &config.AuthSignResp)
	if err != nil {
		return
	}
	return err
}

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

func PushToSingleCid(config ConfigParam, pushReq PushReq) (rsp Resp, err error) {
	client := http.DefaultClient
	payload, _ := json.Marshal(pushReq)
	req, _ := http.NewRequest("POST", config.baseUrl+"/push/single/cid", bytes.NewBuffer(payload))
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

func PushToBatchCid(config ConfigParam, pushReq BatchPush) (rsp Resp, err error) {
	client := http.DefaultClient
	payload, _ := json.Marshal(pushReq)
	req, _ := http.NewRequest("POST", config.baseUrl+"/push/single/batch/cid", bytes.NewBuffer(payload))
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
