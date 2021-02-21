package ebi

import (
	"encoding/json"
	"testing"
)

// TestUnifiedOrderWXNATIVE WX_NATIVE 微信扫码
func TestUnifiedOrderWXNATIVE(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "11111111"
	client.MerchantID = "11111111"
	client.SubAppID = "11111111"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(map[string]string)
	bm["transType"] = "WX_NATIVE"
	bm["transAmt"] = "1"
	// 回调地址
	bm["offlineNotifyUrl"] = "http://127.0.0.1/SMN/src/agency_pay/return.php"
	// 商品ID
	bm["productId"] = "10"
	// 商品名称
	bm["productName"] = "testProuct"
	rsp, err := client.UnifiedOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}

// TestUnifiedOrderALNATIVE AL_NATIVE 支付宝扫码
func TestUnifiedOrderALNATIVE(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "11111111"
	client.MerchantID = "11111111"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(map[string]string)
	bm["transType"] = "AL_NATIVE"
	// 订单金额，单位分
	bm["transAmt"] = "1"
	// 回调地址
	bm["offlineNotifyUrl"] = "http://127.0.0.1/SMN/src/agency_pay/return.php"
	// 商品ID
	bm["productId"] = "10"
	// 商品名称
	bm["productName"] = "testProuct"
	rsp, err := client.UnifiedOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}

// TestUnifiedOrderWXJSAPI 微信 WX_JSAPI
func TestUnifiedOrderWXJSAPI(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "11111111"
	client.MerchantID = "11111111"
	client.SubAppID = "11111111"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(map[string]string)
	bm["transType"] = "WX_JSAPI"
	bm["subAppId"] = "1111111133"
	bm["openid"] = "11111111"
	// 订单金额，单位分
	bm["transAmt"] = "1"
	// 回调地址
	bm["offlineNotifyUrl"] = "http://127.0.0.1/SMN/src/agency_pay/return.php"
	// 商品ID
	bm["productId"] = "10"
	// 商品名称
	bm["productName"] = "testProduct"
	rsp, err := client.UnifiedOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}

func TestUnifiedOrder(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "11111111"
	client.MerchantID = "11111111"
	client.SubAppID = "11111111"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(map[string]string)
	// WX_NATIVE 微信扫码
	// bm["transType"] = "WX_NATIVE"
	// WX_APP 微信 APP 支付
	// bm["transType"] = "WX_APP"
	// CUP_APP 银联控件支付
	// bm["transType"] = "CUP_APP"
	// ALIPAY 支付宝 APP 支付
	// bm["transType"] = "ALIPAY"
	// AL_NATIVE 支付宝扫码
	// bm["transType"] = "AL_NATIVE"
	// ALI_JSAPI 标准支付宝小程序支付
	// bm["transType"] = "ALI_JSAPI"
	// WX_JSAPI 微信 JSAPI
	// bm["transType"] = "WX_JSAPI"
	// 小程序AppID
	// bm["subAppId"] = "wxe000002"
	// 小程序用户OpenID
	// bm["openid"] = "11111111"
	// 订单金额，单位分
	bm["transAmt"] = "1"
	// 回调地址
	bm["offlineNotifyUrl"] = "http://127.0.0.1/SMN/src/agency_pay/return.php"
	// 商品ID
	bm["productId"] = "10"
	// 商品名称
	bm["productName"] = "f"
	rsp, err := client.UnifiedOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}

func TestQueryOrder(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "11111111"
	client.MerchantID = "11111111"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(map[string]string)
	// 订单ID
	bm["orderId"] = "202102201638282523"
	// 请求号ID
	bm["requestId"] = "202102201638282524"
	rsp, err := client.QueryOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}

func TestOrder(t *testing.T) {
	Order()
}
