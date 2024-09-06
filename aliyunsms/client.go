// Package aliyunsms 阿里云SMS短信服务
package aliyunsms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type resp struct {
	BizId     string
	Code      string
	Message   string
	RequestId string
}
type Client struct {
	endPoint        string
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
	PhoneNumbers    string
	Content         string
	paras           map[string]string
	resp            resp
}

// NewClient 初始化阿里云短信服务客户端
//
// accessKeyID 访问ID
//
// accessKeySecret 访问密钥
func NewClient(accessKeyID, accessKeySecret string) (client *Client) {
	paras := make(map[string]string)
	paras["SignatureMethod"] = "HMAC-SHA1"
	paras["AccessKeyId"] = accessKeyID
	paras["SignatureVersion"] = "1.0"
	paras["Format"] = "JSON"
	paras["Action"] = "SendSms"
	paras["Version"] = "2017-05-25"
	paras["RegionId"] = "cn-hangzhou"
	return &Client{
		endPoint:        "dysmsapi.aliyuncs.com",
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		paras:           paras}
}

// SendSms 发送短信
// phoneNumbers 接收短信的手机号码
//
// signName 短信签名名称
//
// templateCode 短信模板ID
//
// templateParam 短信模板变量对应的实际值,JSON格式
//
// 返回参数 saveURL 成功上传后保存路径
func (s *Client) SendSms(phoneNumbers, signName, templateCode, templateParam string) (err error) {
	timezone, err := time.LoadLocation("GMT0")
	if err != nil {
		return
	}
	s.paras["PhoneNumbers"] = phoneNumbers
	s.paras["TemplateCode"] = templateCode
	s.paras["TemplateParam"] = templateParam
	s.paras["SignName"] = signName
	s.paras["SignatureNonce"] = uuid.NewV4().String()
	s.paras["Timestamp"] = time.Now().In(timezone).Format("2006-01-02T15:04:05Z")
	delete(s.paras, "Signature")
	parasIdx := make([]string, 0)
	for k := range s.paras {
		parasIdx = append(parasIdx, k)
	}
	sort.Strings(parasIdx)
	sortedQueryString := ""
	for _, v := range parasIdx {
		sortedQueryString = sortedQueryString + "&" + specialUrlEncode(v) + "=" + specialUrlEncode(s.paras[v])
	}
	sortedQueryString = sortedQueryString[1:]
	stringToSign := "GET" + "&" + specialUrlEncode("/") + "&" + specialUrlEncode(sortedQueryString)
	signStr := sign(s.AccessKeySecret+"&", stringToSign)
	signature := specialUrlEncode(signStr)
	urlStr := "http://" + s.endPoint + "/?Signature=" + signature + "&" + sortedQueryString
	resp, err := http.Get(urlStr)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	json.Unmarshal(bodyContent, &s.resp)
	if s.resp.Code != "OK" {
		return errors.New(s.resp.Code)
	}
	return nil
}

func specialUrlEncode(value string) string {
	result := url.QueryEscape(value)
	result = strings.Replace(result, "+", "%20", -1)
	result = strings.Replace(result, "*", "%2A", -1)
	result = strings.Replace(result, "%7E", "~", -1)
	return result
}

func sign(accessSecret, strToSign string) string {
	mac := hmac.New(sha1.New, []byte(accessSecret))
	mac.Write([]byte(strToSign))
	signData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signData)
}
