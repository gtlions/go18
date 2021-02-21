package ebi

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/gogap/mahonia"
	"golang.org/x/crypto/pkcs12"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Client struct {
	init         bool   `json:"init,omitempty"`
	TimeOut      int    `json:"time_out,omitempty"`
	Pfx          string `json:"pfx,omitempty"`
	PfxPasswd    string `json:"pfxPasswd,omitempty"`
	Charset      string `json:"charset,omitempty"`
	Version      string `json:"version,omitempty"`
	SignType     string `json:"signType,omitempty"`
	MerchantSign string `json:"merchantSign,omitempty"`
	MerchantCert string `json:"merchantCert,omitempty"`
	Service      string `json:"service,omitempty"`
	MerchantID   string `json:"merchantId,omitempty"`
	SubAppID     string `json:"subAppId,omitempty"`
	Currency     string `json:"currency,omitempty"`
}

// Init 电银支付客户端初始化
//
func (c *Client) Init() (err error) {
	if c.MerchantID == "" {
		return fmt.Errorf("缺少初始化参数: [ MerchantID ]")
	}
	if c.Pfx == "" {
		return fmt.Errorf("缺少初始化参数: [ Pfx ]")
	}
	if c.PfxPasswd == "" {
		return fmt.Errorf("缺少初始化参数: [ PfxPasswd ]")
	}
	if c.init {
		return nil
	}
	c.init = true
	return nil
}

// isInit 是否成功初始化
//
func (c *Client) isInit() (bool, error) {
	if !c.init {
		return c.init, fmt.Errorf("未初始化客户端")
	}
	return c.init, nil
}

// sign 电银支付请求数据加密、签名
//
// data 请求参数
//
func (c *Client) sign(data map[string]string) (map[string]string, error) {
	var (
		err           error
		encryptedData []byte
		privateKeys   []interface{}
		certificates  []*x509.Certificate
	)
	s := make([]string, len(data))
	for k := range data {
		s = append(s, k)
	}
	sort.Strings(s)
	str := ""
	enc := mahonia.NewEncoder("GBK")
	for _, v := range s {
		if data[v] == "" || v == "serverSign" || v == "serverCert" {
			continue
		}
		if str != "" {
			str += "&"
		}
		str += v + "=" + enc.ConvertString(data[v])
	}
	if privateKeys, certificates, err = pkcs12DecodeAll(c.Pfx, c.PfxPasswd); err != nil {
		return nil, err
	}
	hash := sha256.New()
	hash.Write([]byte(str))
	if encryptedData, err = rsa.SignPKCS1v15(rand.Reader, privateKeys[0].(*rsa.PrivateKey), crypto.SHA256, hash.Sum(nil)); err != nil {
		return nil, err
	}
	data["merchantSign"] = strings.ToUpper(hex.EncodeToString(encryptedData))
	data["merchantCert"] = strings.ToUpper(hex.EncodeToString(certificates[0].Raw))
	return data, nil
}

type EbiResponse struct {
	RspCode         string    `json:"rspCode,omitempty"`
	RspMessage      string    `json:"rspMessage,omitempty"`
	Charset         string    `json:"charset,omitempty"`
	Version         string    `json:"version,omitempty"`
	SignType        string    `json:"signType,omitempty"`
	ServerCert      string    `json:"serverCert,omitempty"`
	ServerSign      string    `json:"serverSign,omitempty"`
	Service         string    `json:"service,omitempty"`
	TradeType       string    `json:"tradeType,omitempty"`
	MerchantID      string    `json:"merchantId,omitempty"`
	OrderID         string    `json:"orderId,omitempty"`
	RequestID       string    `json:"requestId,omitempty"`
	OrderTime       string    `json:"orderTime,omitempty"`
	TradeNO         string    `json:"tradeNo,omitempty"`
	TransAmt        string    `json:"transAmt,omitempty"`
	TransState      string    `json:"transState,omitempty"`
	PayTime         string    `json:"payTime,omitempty"`
	PromotionDetail string    `json:"promotionDetail,omitempty"`
	SettleDate      string    `json:"settleDate,omitempty"`
	SettleTransAmt  string    `json:"settleTransAmt,omitempty"`
	ChannelNo       string    `json:"channelNo,omitempty"`
	CertID          string    `json:"certId,omitempty"`
	FeeAmt          string    `json:"feeAmt,omitempty"`
	PayUrl          string    `json:"pay_url,omitempty"`
	PayInfo         string    `json:"payInfo,omitempty"`
	WcPayData       WcPayData `json:"wcPayData,omitempty"`
	Tn              string    `json:"tn,omitempty"`
	ExtendInfo      string    `json:"extendInfo,omitempty"`
	JsAppID         string    `json:"jsAppId,omitempty"`
	JsAppUrl        string    `json:"jsAppUrl,omitempty"`
}

type WcPayData struct {
	AppID     string `json:"appId,omitempty"`
	PartnerID string `json:"partnerId,omitempty"`
	PrepayID  string `json:"prepayId,omitempty"`
	Package   string `json:"package,omitempty"`
	NonceStr  string `json:"nonceStr,omitempty"`
	TimeStamp string `json:"timeStamp,omitempty"`
	PaySign   string `json:"paySign,omitempty"`
	SignType  string `json:"signType,omitempty"`
}

// pkcs12DecodeAll 解析pfx证书
//
// pfxFileName 文件路径
//
// password 访问密码
//
func pkcs12DecodeAll(pfxFileName string, password string) ([]interface{}, []*x509.Certificate, error) {
	pfxData := new([]byte)
	if keyFile, err := ioutil.ReadFile(pfxFileName); err != nil {
		return nil, nil, err
	} else {
		pfxData = &keyFile
	}
	var privateKeys []interface{}
	var certificates []*x509.Certificate
	blocks, err := pkcs12.ToPEM(*pfxData, password)
	if err != nil {
		return nil, nil, err
	}
	for _, b := range blocks {
		if b.Type == "CERTIFICATE" {
			certs, err := x509.ParseCertificates(b.Bytes)
			if err != nil {
				return nil, nil, err
			}
			certificates = append(certificates, certs...)

		} else if b.Type == "PRIVATE KEY" {
			privateKey, err := x509.ParsePKCS1PrivateKey(b.Bytes)
			if err != nil {
				return nil, nil, err
			}
			privateKeys = append(privateKeys, privateKey)
		}
	}
	return privateKeys, certificates, err
}

func gbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
