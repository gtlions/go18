package aliyunoss

import (
	"testing"
)

func TestPutFromFile(t *testing.T) {
	endPoint := "endPoint"
	accessKeyID := "accessKeyID"
	accessKeySecret := "accessKeySecret"
	bucket := "bucket"
	object := "object"
	client := NewClient(endPoint, accessKeyID, accessKeySecret, bucket, object)
	saveURL, err := client.PutFromFile("aliyunoss_test.txt", "filename1")
	if err != nil {
		t.Fatal("操作失败", err)
	}
	t.Logf("操作成功,保存至: %s", saveURL)
}

func TestDeleteFile(t *testing.T) {
	endPoint := "endPoint"
	accessKeyID := "accessKeyID"
	accessKeySecret := "accessKeySecret"
	bucket := "bucket"
	object := "object"
	client := NewClient(endPoint, accessKeyID, accessKeySecret, bucket, object)
	err := client.DeleteFile("aliyunoss_test.txt")
	if err != nil {
		t.Fatal("操作失败", err)
	}
	t.Logf("操作成功")
}
