// Package aliyunoss 阿里云SMS短信服务
package aliyunsms

import (
	"fmt"
	"testing"
)

func TestSendSms(t *testing.T) {
	accessKeyID := "LTAI4G4kHXWANWNkv6xdVEfq"
	accessKeySecret := "dYBYBMGdnjUeRdAeuWM7YCcjgVXUni"
	client := NewClient(accessKeyID, accessKeySecret)
	err := client.SendSms("13774553612", "芯美丽", "SMS_204756660", fmt.Sprintf(`{"code":%s}`, "123456"))

	if err != nil {
		t.Fatal("操作失败", err)
	}
	t.Logf("操作成功")
}
