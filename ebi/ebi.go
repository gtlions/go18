package ebi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/schema"
	"github.com/gtlions/gos10i"
)

const BASEURL = "http://pay.uat.chinaebi.com:50080/mrpos/cashier"

// UnifiedOrder 支付统一下单
//
// bm 提交的数据
//
func (c *Client) UnifiedOrder(bm map[string]string) (rsp EbiResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	if bm["transAmt"] == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "transAmt")
	} else {
		if transAmt, err := strconv.Atoi(bm["transAmt"]); err != nil {
			return rsp, fmt.Errorf("参数错误: [ %s ]", "transAmt")
		} else if transAmt <= 0 {
			return rsp, fmt.Errorf("金额必须大于0: [ %s ]", "transAmt")
		}
	}
	if bm["transType"] == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "transType")
	}
	if bm["offlineNotifyUrl"] == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "offlineNotifyUrl")
	}
	if bm["productId"] == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "productId")
	}
	if bm["productName"] == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "productName")
	}
	if bm["productDesc"] == "" {
		bm["productDesc"] = bm["productName"]
	}
	if bm["charset"] == "" {
		bm["charset"] = "00"
	}
	if bm["version"] == "" {
		bm["version"] = "1.3"
	}
	if bm["signType"] == "" {
		bm["signType"] = "RSA"
	}
	if bm["service"] == "" {
		bm["service"] = "DowDirectPay"
	}
	nonceStr := gos10i.XOrderNoFromNow()
	if bm["orderId"] == "" {
		bm["orderId"] = nonceStr
	}
	if bm["orderTime"] == "" {
		bm["orderTime"] = time.Now().Format("20060102")
	}
	if bm["transAmt"] == "" {
		bm["transAmt"] = "111"
	}
	if bm["currency"] == "" {
		bm["currency"] = "CNY"
	}
	if bm["validUnit"] == "" {
		bm["validUnit"] = "00"
	}
	if bm["validNum"] == "" {
		bm["validNum"] = "15"
	}
	if bm["requestId"] == "" {
		bm["requestId"] = nonceStr
	}
	if bm["clientIP"] == "" {
		bm["clientIP"] = "127.0.0.1"
	}
	if bm["accFlag"] == "" {
		bm["accFlag"] = "0"
	}
	bm["merchantId"] = c.MerchantID
	if bm["subAppId"] == "" && bm["transType"] == "WX_JSAPI" {
		if c.SubAppID == "" {
			return rsp, fmt.Errorf("缺少必要参数 [ subAppId ]")
		}
		bm["subAppId"] = c.SubAppID
	}
	bm, err = c.sign(bm)
	if err != nil {
		return rsp, err
	}
	urlValues := url.Values{}
	for k, v := range bm {
		urlValues.Add(k, v)
	}
	client := http.DefaultClient
	if c.TimeOut > 0 {
		client.Timeout = time.Second * time.Duration(c.TimeOut)
	} else {
		client.Timeout = time.Second * 15
	}
	req, _ := http.NewRequest("POST", BASEURL, strings.NewReader(urlValues.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return rsp, err
	}
	defer resp.Body.Close()
	rspBody, _ := ioutil.ReadAll(resp.Body)
	s := ""
	utf8, err := gbkToUtf8(rspBody)
	if err != nil {
		return rsp, err
	} else {
		s = string(utf8)
	}
	urlRsp := "http://127.0.0.1?" + s
	qryRsp, err := url.Parse(urlRsp)
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if err = json.Unmarshal(rspBody, &rsp); err != nil {
			return rsp, err
		}
	}
	rsp.Charset = ""
	rsp.Version = ""
	rsp.SignType = ""
	rsp.ServerCert = ""
	rsp.ServerSign = ""
	rsp.Service = ""
	rsp.OrderID = bm["orderId"]
	rsp.RequestID = bm["requestId"]
	return rsp, err
}

// QueryOrder 支付订单查询
//
// bm 提交的数据
//
func (c *Client) QueryOrder(bm map[string]string) (rsp EbiResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	if bm["orderId"] == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ orderId ]")
	}
	if bm["requestId"] == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ requestId ]")
	}
	if bm["charset"] == "" {
		bm["charset"] = "00"
	}
	if bm["version"] == "" {
		bm["version"] = "1.1"
	}
	if bm["signType"] == "" {
		bm["signType"] = "RSA"
	}
	if bm["service"] == "" {
		bm["service"] = "DirectOrderSearch"
	}
	bm["merchantId"] = c.MerchantID
	bm, err = c.sign(bm)
	if err != nil {
		return rsp, err
	}
	client := &http.Client{}
	urlValues := url.Values{}
	for k, v := range bm {
		urlValues.Add(k, v)
	}
	req, _ := client.PostForm(BASEURL, urlValues)
	rspBody, _ := ioutil.ReadAll(req.Body)
	s := ""
	utf8, err := gbkToUtf8(rspBody)
	if err != nil {
		return rsp, err
	} else {
		s = string(utf8)
	}
	urlRsp := "http://127.0.0.1?" + s
	qryRsp, err := url.Parse(urlRsp)
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if err = json.Unmarshal(rspBody, &rsp); err != nil {
			return rsp, err
		}
	}
	rsp.Charset = ""
	rsp.Version = ""
	rsp.SignType = ""
	rsp.ServerCert = ""
	rsp.ServerSign = ""
	rsp.Service = ""
	return rsp, err
}
