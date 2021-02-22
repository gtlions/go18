package ebi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gogap/mahonia"
	"github.com/gorilla/schema"
	"github.com/gtlions/gos10i"
)

const BASEURL = "http://pay.uat.chinaebi.com:50080/mrpos/cashier"
const TIMEOUT = 15

// UnifiedOrder 支付统一下单
//
// bm 提交的数据
//
func (c *Client) UnifiedOrder(bm BodyMap) (rsp UnifiedOrderResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	if bm.Get("transAmt") == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "transAmt")
	} else {
		if transAmt, err := strconv.Atoi(bm.Get("transAmt")); err != nil {
			return rsp, fmt.Errorf("参数错误: [ %s ]", "transAmt")
		} else if transAmt <= 0 {
			return rsp, fmt.Errorf("金额必须大于0: [ %s ]", "transAmt")
		}
	}
	if bm.Get("transType") == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "transType")
	}
	if bm.Get("offlineNotifyUrl") == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "offlineNotifyUrl")
	}
	if bm.Get("productId") == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "productId")
	}
	if bm.Get("productName") == "" {
		return rsp, fmt.Errorf("缺少必要参数: [ %s ]", "productName")
	}
	// if bm.Get("productDesc") == "" {
	// 	bm.Set("productDesc", bm.Get("productName"))
	// }
	if bm.Get("charset") == "" {
		bm.Set("charset", "00")
	}
	if bm.Get("version") == "" {
		bm.Set("version", "1.3")
	}
	if bm.Get("signType") == "" {
		bm.Set("signType", "RSA")
	}
	if bm.Get("service") == "" {
		bm.Set("service", "DowDirectPay")
	}
	nonceStr := gos10i.XOrderNoFromNow()
	if bm.Get("orderId") == "" {
		bm.Set("orderId", nonceStr)
	}
	if bm.Get("orderTime") == "" {
		bm.Set("orderTime", time.Now().Format("20060102"))
	}
	if bm.Get("transAmt") == "" {
		bm.Set("transAmt", "111")
	}
	if bm.Get("currency") == "" {
		bm.Set("currency", "CNY")
	}
	if bm.Get("validUnit") == "" {
		bm.Set("validUnit", "00")
	}
	if bm.Get("validNum") == "" {
		bm.Set("validNum", "15")
	}
	if bm.Get("requestId") == "" {
		bm.Set("requestId", nonceStr)
	}
	if bm.Get("clientIP") == "" {
		bm.Set("clientIP", "127.0.0.1")
	}
	if bm.Get("accFlag") == "" {
		bm.Set("accFlag", "0")
	}
	bm.Set("merchantId", c.MerchantID)
	if bm.Get("subAppId") == "" && bm.Get("transType") == "WX_JSAPI" {
		if c.SubAppID == "" {
			return rsp, fmt.Errorf("缺少必要参数 [ subAppId ]")
		}
		bm.Set("subAppId", c.SubAppID)
	}
	bm, err = c.sign(bm)
	if err != nil {
		return rsp, err
	}
	enc := mahonia.NewEncoder("GBK")
	urlValues := url.Values{}
	for k := range bm {
		urlValues.Add(k, enc.ConvertString(bm.Get(k)))
	}
	client := http.DefaultClient
	if c.TimeOut > 0 {
		client.Timeout = time.Second * time.Duration(c.TimeOut)
	} else {
		client.Timeout = time.Second * TIMEOUT
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
	utf8, err := gos10i.GbkToUtf8(rspBody)
	if err != nil {
		return rsp, err
	} else {
		s = string(utf8)
	}
	urlRsp := "http://127.0.0.1?" + s
	qryRsp, err := url.Parse(urlRsp)
	if err != nil {
		return rsp, err
	}
	if c.Debug {
		log.Println("qryRsp.Query:", qryRsp.Query())
	}
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if c.Debug {
			log.Println("schema.NewDecoder->err:", err)
		}
		if err = json.Unmarshal(rspBody, &rsp); err != nil {
			return rsp, err
		}
	}
	if rsp.WcPayData != "" {
		if err = json.Unmarshal([]byte(rsp.WcPayData), &rsp.WxPayData); err != nil {
			return rsp, err
		}
		rsp.WcPayData = ""
	}
	rsp.Charset = ""
	rsp.Version = ""
	rsp.SignType = ""
	rsp.ServerCert = ""
	rsp.ServerSign = ""
	rsp.Service = ""
	rsp.OrderID = bm.Get("orderId")
	rsp.RequestID = bm.Get("requestId")
	return rsp, err
}

// QueryOrder 支付订单查询
//
// bm 提交的数据
//
func (c *Client) QueryOrder(bm BodyMap) (rsp QueryOrderResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	if bm.Get("orderId") == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ orderId ]")
	}
	if bm.Get("requestId") == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ requestId ]")
	}
	if bm.Get("charset") == "" {
		bm.Set("charset", "00")
	}
	if bm.Get("version") == "" {
		bm.Set("version", "1.1")
	}
	if bm.Get("signType") == "" {
		bm.Set("signType", "RSA")
	}
	if bm.Get("service") == "" {
		bm.Set("service", "DirectOrderSearch")
	}
	bm.Set("merchantId", c.MerchantID)
	bm, err = c.sign(bm)
	if err != nil {
		return rsp, err
	}
	client := &http.Client{}
	urlValues := url.Values{}
	for k := range bm {
		urlValues.Add(k, bm.Get(k))
	}
	req, _ := client.PostForm(BASEURL, urlValues)
	rspBody, _ := ioutil.ReadAll(req.Body)
	s := ""
	utf8, err := gos10i.GbkToUtf8(rspBody)
	if err != nil {
		return rsp, err
	} else {
		s = string(utf8)
	}
	urlRsp := "http://127.0.0.1?" + s
	qryRsp, err := url.Parse(urlRsp)
	if err != nil {
		return rsp, err
	}
	if c.Debug {
		log.Println("qryRsp.Query:", qryRsp.Query())
	}
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if c.Debug {
			log.Println("schema.NewDecoder->err:", err)
		}
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

// UnifiedOrderDiscard 支付统一下单
//
// bm 提交的数据
//
func (c *Client) UnifiedOrderDiscard(bm map[string]string) (rsp UnifiedOrderResponse, err error) {
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
	bm, err = c.signDiscard(bm)
	if err != nil {
		return rsp, err
	}
	enc := mahonia.NewEncoder("GBK")
	urlValues := url.Values{}
	for k, v := range bm {
		urlValues.Add(k, enc.ConvertString(v))
	}
	client := http.DefaultClient
	if c.TimeOut > 0 {
		client.Timeout = time.Second * time.Duration(c.TimeOut)
	} else {
		client.Timeout = time.Second * 30
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
	utf8, err := gos10i.GbkToUtf8(rspBody)
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

// QueryOrderDiscard 支付订单查询
//
// bm 提交的数据
//
func (c *Client) QueryOrderDiscard(bm map[string]string) (rsp UnifiedOrderResponse, err error) {
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
	bm, err = c.signDiscard(bm)
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
	utf8, err := gos10i.GbkToUtf8(rspBody)
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
