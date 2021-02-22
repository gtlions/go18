package ebi

import (
	"encoding/json"
	"testing"
)

// TestUnifiedOrder
func TestUnifiedOrder(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "PfxPasswd"
	client.MerchantID = "MerchantID"
	client.SubAppID = "SubAppID"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(BodyMap)
	// bm.Set("transType", "WX_NATIVE")
	// bm.Set("transType", "AL_NATIVE")
	bm.Set("transType", "WX_JSAPI")
	bm.Set("openid", "openid")

	bm.Set("transAmt", "102111")
	bm.Set("offlineNotifyUrl", "http://127.0.0.1/SMN/src/agency_pay/return.php")
	bm.Set("productId", "10")
	bm.Set("productName", "testProuct商品名称")
	rsp, err := client.UnifiedOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}

// TestQueryOrder
func TestQueryOrder(t *testing.T) {
	client := Client{}
	client.Pfx = "merchant.pfx"
	client.PfxPasswd = "PfxPasswd"
	client.MerchantID = "MerchantID"
	err := client.Init()
	if err != nil {
		t.Error(err)
	}
	bm := make(BodyMap)
	bm.Set("orderId", "orderId")
	bm.Set("requestId", "requestId")
	rsp, err := client.QueryOrder(bm)
	if err != nil {
		t.Error(err)
	}
	r, _ := json.MarshalIndent(rsp, "", "	")
	t.Logf("\n\n%s\n\n", string(r))
}
