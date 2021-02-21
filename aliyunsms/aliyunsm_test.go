// Package aliyunoss 阿里云SMS短信服务
package aliyunsms

import (
	"fmt"
	"testing"
)

func TestSendSms(t *testing.T) {
	accessKeyID := "accessKeyID"
	accessKeySecret := "accessKeySecret"
	client := NewClient(accessKeyID, accessKeySecret)
	err := client.SendSms("137745536xx", "xxx", "SMS_xxxxxx", fmt.Sprintf(`{"code":%s}`, "123456"))

	if err != nil {
		t.Fatal("操作失败", err)
	}
	t.Logf("操作成功")
}
