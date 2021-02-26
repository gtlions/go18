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

// UnifiedOrder 统一下单
//
// bm 提交的数据
//
func (c *Client) UnifiedOrder(bm BodyMap) (rsp UnifiedOrderResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	nonceStr := gos10i.XOrderNoFromNow()
	bm.Set("merchantId", c.MerchantID)
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
	if c.Debug {
		s, err := json.MarshalIndent(bm, "", "	")
		if err != nil {
			log.Println("MarshalIndent Request urlValues err:", err)
		}
		log.Println("Request Values:", string(s))
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
		s, err := json.MarshalIndent(qryRsp.Query(), "", "	")
		if err != nil {
			log.Println("MarshalIndent Response Values err:", err)
		}
		log.Println("Response Values:", string(s))
	}
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if c.Debug {
			log.Println("Decoder Response Values Error:", err)
		}
		if err = json.Unmarshal(rspBody, &rsp); err != nil {
			return rsp, err
		}
	}
	if rsp.WcPayData != "" {
		if err = json.Unmarshal([]byte(rsp.WcPayData), &rsp.WechatPayData); err != nil {
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

// QueryOrder 订单查询
//
// bm 提交的数据
//
func (c *Client) QueryOrder(bm BodyMap) (rsp QueryOrderResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	bm.Set("merchantId", c.MerchantID)
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

	bm, err = c.sign(bm)
	if err != nil {
		return rsp, err
	}
	client := &http.Client{}
	urlValues := url.Values{}
	for k := range bm {
		urlValues.Add(k, bm.Get(k))
	}
	if c.Debug {
		s, err := json.MarshalIndent(bm, "", "	")
		if err != nil {
			log.Println("MarshalIndent Request urlValues err:", err)
		}
		log.Println("Request Values:", string(s))
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
		s, err := json.MarshalIndent(qryRsp.Query(), "", "	")
		if err != nil {
			log.Println("MarshalIndent Response Values err:", err)
		}
		log.Println("Response Values:", string(s))
	}
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if c.Debug {
			log.Println("Decoder Response Values Error:", err)
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

// RefundOrder 订单退款
//
// bm 提交的数据
//
func (c *Client) RefundOrder(bm BodyMap) (rsp RefundOrderResponse, err error) {
	if _, err := c.isInit(); err != nil {
		return rsp, err
	}
	nonceStr := gos10i.XOrderNoFromNow()
	bm.Set("merchantId", c.MerchantID)
	if bm.Get("charset") == "" {
		bm.Set("charset", "00")
	}
	if bm.Get("version") == "" {
		bm.Set("version", "1.0")
	}
	if bm.Get("signType") == "" {
		bm.Set("signType", "RSA")
	}
	if bm.Get("service") == "" {
		bm.Set("service", "DYRefund")
	}
	if bm.Get("orderId") == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ orderId ]")
	}
	if bm.Get("refundAmount") == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ refundAmount ]")
	}
	if bm.Get("requestId") == "" {
		bm.Set("requestId", nonceStr)
	}
	if bm.Get("refundId") == "" {
		bm.Set("refundId", nonceStr)
	}
	if bm.Get("offlineNotifyUrl") == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ offlineNotifyUrl ]")
	}
	if bm.Get("clientIP") == "" {
		bm.Set("clientIP", "127.0.0.1")
	}
	if bm.Get("requestId") == "" {
		return rsp, fmt.Errorf("缺少必要参数 [ requestId ]")
	}
	bm, err = c.sign(bm)
	if err != nil {
		return rsp, err
	}
	client := &http.Client{}
	urlValues := url.Values{}
	for k := range bm {
		urlValues.Add(k, bm.Get(k))
	}
	if c.Debug {
		s, err := json.MarshalIndent(bm, "", "	")
		if err != nil {
			log.Println("MarshalIndent Request urlValues err:", err)
		}
		log.Println("Request Values:", string(s))
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
		s, err := json.MarshalIndent(qryRsp.Query(), "", "	")
		if err != nil {
			log.Println("MarshalIndent Response Values err:", err)
		}
		log.Println("Response Values:", string(s))
	}
	var decoder = schema.NewDecoder()
	err = decoder.Decode(&rsp, qryRsp.Query())
	if err != nil {
		// return rsp, err
		if c.Debug {
			log.Println("Decoder Response Values Error:", err)
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

// ParseNotify 解析支付异步通知的参数
//
// req：*http.Request
//
// 返回参数 rsp 请求的参数
//
// 返回参数 err 错误信息
//
func ParseNotify(req *http.Request) (rsp *NotifyRequest, err error) {
	var decoder = schema.NewDecoder()
	rspTemp := NotifyRequest{}
	if err != nil {
		return rsp, err
	}
	err = decoder.Decode(&rspTemp, req.PostForm)
	if err != nil {
		return rsp, err
	}
	rsp = &rspTemp
	return rsp, err
}
