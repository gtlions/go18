package ebi

import (
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
)

type Client struct {
	Debug bool `json:"debug"`
	init  bool `json:"-"`
	// 超时
	TimeOut int `json:"time_out,omitempty"`
	// 服务器证书文件
	Pfx string `json:"pfx,omitempty"`
	// 服务器证书密码
	PfxPasswd string `json:"pfxPasswd,omitempty"`
	// 字符集
	Charset string `json:"charset,omitempty"`
	// 版本
	Version string `json:"version,omitempty"`
	// 签名方式
	SignType string `json:"signType,omitempty"`
	// 服务器签名
	MerchantSign string `json:"merchantSign,omitempty"`
	// 服务器证书
	MerchantCert string `json:"merchantCert,omitempty"`
	// 交易接口
	Service string `json:"service,omitempty"`
	// 商户号
	MerchantID string `json:"merchantId,omitempty"`
	// 子商户公众账号ID
	SubAppID string `json:"subAppId,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty"`
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
func (c *Client) sign(bm BodyMap) (BodyMap, error) {
	var (
		err           error
		encryptedData []byte
		privateKeys   []interface{}
		certificates  []*x509.Certificate
	)
	s := make([]string, len(bm))
	for k := range bm {
		s = append(s, k)
	}
	sort.Strings(s)
	str := ""
	enc := mahonia.NewEncoder("GBK")
	for _, v := range s {
		if bm.Get(v) == "" || v == "serverSign" || v == "serverCert" {
			continue
		}
		if str != "" {
			str += "&"
		}
		str += v + "=" + enc.ConvertString(bm.Get(v))
	}
	if privateKeys, certificates, err = pkcs12DecodeAll(c.Pfx, c.PfxPasswd); err != nil {
		return nil, err
	}
	hash := sha256.New()
	hash.Write([]byte(str))
	if encryptedData, err = rsa.SignPKCS1v15(rand.Reader, privateKeys[0].(*rsa.PrivateKey), crypto.SHA256, hash.Sum(nil)); err != nil {
		return nil, err
	}
	bm["merchantSign"] = strings.ToUpper(hex.EncodeToString(encryptedData))
	bm["merchantCert"] = strings.ToUpper(hex.EncodeToString(certificates[0].Raw))
	return bm, nil
}

// signDiscard 电银支付请求数据加密、签名
//
// data 请求参数
//
func (c *Client) signDiscard(data map[string]string) (map[string]string, error) {
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

// UnifiedOrderResponse 订单接口请求响应
type UnifiedOrderResponse struct {
	// 返回状态码
	RspCode string `json:"rspCode,omitempty"`
	// 返回信息
	RspMessage string `json:"rspMessage,omitempty"`
	// 字符集
	Charset string `json:"charset,omitempty"`
	// 接口版本
	Version string `json:"version,omitempty"`
	// 签名类型
	SignType string `json:"signType,omitempty"`
	// 服务器证书
	ServerCert string `json:"serverCert,omitempty"`
	// 服务器签名
	ServerSign string `json:"serverSign,omitempty"`
	// 交易接口
	Service string `json:"service,omitempty"`
	// 交易类型
	TradeType string `json:"tradeType,omitempty"`
	// 商户号
	MerchantID string `json:"merchantId,omitempty"`
	// 商户订单号
	OrderID string `json:"orderId,omitempty"`
	// 订单日期
	OrderTime string `json:"orderTime,omitempty"`
	// 请求号
	RequestID string `json:"requestId,omitempty"`
	// 支付流水号
	TradeNO string `json:"tradeNo,omitempty"`
	// 交易金额
	TransAmt string `json:"transAmt,omitempty"`
	// 支付二维码连接
	PayInfo string `json:"payInfo,omitempty"`
	// 微信支付返回数据
	WcPayData string `json:"wcPayData,omitempty"`
	// 微信支付返回数据
	WechatPayData *WechatPayData `json:"wechatPayData,omitempty"`
	// 银联流水号
	Tn string `json:"tn,omitempty"`
	// 证书序列号
	CertID string `json:"certId,omitempty"`
	// 扩展信息
	ExtendInfo string `json:"extendInfo,omitempty"`
	// 小程序ID
	JsAppID string `json:"jsAppId,omitempty"`
	// 小程序地址
	JsAppUrl string `json:"jsAppUrl,omitempty"`
	PayUrl   string `json:"payUrl,omitempty"`
}

// WechatPayData 微信支付信息响应
type WechatPayData struct {
	// 应用ID
	AppID string `json:"appId,omitempty"`
	// 从业机构号
	PartnerID string `json:"partnerId,omitempty"`
	// 预支付交易会话ID
	PrepayID string `json:"prepayId,omitempty"`
	// 订单详情扩展字符串
	Package string `json:"package,omitempty"`
	// 随机字符串
	NonceStr string `json:"nonceStr,omitempty"`
	// 时间戳
	TimeStamp string `json:"timeStamp,omitempty"`
	// 签名
	PaySign string `json:"paySign,omitempty"`
	// 签名方式
	SignType string `json:"signType,omitempty"`
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

// QueryOrderResponse 订单查询接口响应
type QueryOrderResponse struct {
	// 返回状态码
	RspCode string `json:"rspCode,omitempty"`
	// 返回信息
	RspMessage string `json:"rspMessage,omitempty"`
	// 字符集
	Charset string `json:"charset,omitempty"`
	// 接口版本
	Version string `json:"version,omitempty"`
	// 签名类型
	SignType string `json:"signType,omitempty"`
	// 服务器证书
	ServerCert string `json:"serverCert,omitempty"`
	// 服务器签名
	ServerSign string `json:"serverSign,omitempty"`
	// 交易接口
	Service string `json:"service,omitempty"`
	// 商户订单号
	OrderID string `json:"orderId,omitempty"`
	// 订单日期
	OrderTime string `json:"orderTime,omitempty"`
	// 支付流水号
	TradeNO string `json:"tradeNo,omitempty"`
	// 商户号
	MerchantID string `json:"merchantId,omitempty"`
	// 交易金额
	TransAmt string `json:"transAmt,omitempty"`
	// 交易状态 S-成功,P-交易失败,P-交易处理中
	TransState string `json:"transState,omitempty"`
	// 支付完成时间
	PayTime string `json:"payTime,omitempty"`
	// 清算日期
	SettleDate string `json:"settleDate,omitempty"`
	// 签名
	SettleTransAmt string `json:"settleTransAmt,omitempty"`
	// 通道流水号
	ChannelNo string `json:"channelNo,omitempty"`
	// 证书序列号
	CertID string `json:"certId,omitempty"`
	// 扩展信息
	ExtendInfo string `json:"extendInfo,omitempty"`
	// 交易手续费
	FeeAmt string `json:"feeAmt,omitempty"`
	//
	PromotionDetail string `json:"promotionDetail,omitempty"`
}

// 支付异步通知的参数
type NotifyRequest struct {
	// 字符集
	Charset string `json:"charset,omitempty"`
	// 接口版本
	Version string `json:"version,omitempty"`
	// 签名类型
	SignType string `json:"signType,omitempty"`
	// 服务器证书
	ServerCert string `json:"serverCert,omitempty"`
	// 服务器签名
	ServerSign string `json:"serverSign,omitempty"`
	// 商户订单号
	OrderID string `json:"orderId,omitempty"`
	// 支付流水号
	TradeNO string `json:"tradeNo,omitempty"`
	// 商户号
	MerchantID string `json:"merchantId,omitempty"`
	// 交易金额
	TransAmt string `json:"transAmt,omitempty"`
	// 交易状态 S-成功,P-交易失败,P-交易处理中
	TransState string `json:"transState,omitempty"`
	// 支付完成时间
	PayTime string `json:"payTime,omitempty"`
	// 商户私有域
	BackParam string `json:"backParam,omitempty"`
	// 通道流水号
	ChannelNo string `json:"channelNo,omitempty"`
	// 清算日期
	SettleDate string `json:"settleDate,omitempty"`
	// 证书序列号
	CertID string `json:"certId,omitempty"`
	// 扩展信息
	ExtendInfo string `json:"extendInfo,omitempty"`
}
