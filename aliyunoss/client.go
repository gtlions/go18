// Package aliyunoss 阿里云OSS相关功能
package aliyunoss

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Object          string
	ossClient       *oss.Client
}

// 初始化阿里云OSS客户端
//
// endPoint: 访问域名
//
// accessKeyID: 访问ID
//
// accessKeySecret: 访问密钥
//
// bucket: 存储空间
//
// object: 对象
func NewClient(endPoint, accessKeyID, accessKeySecret, bucket, object string) (client *Client) {
	ossClient, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	if err != nil {
		return
	}
	return &Client{
		EndPoint:        endPoint,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		Bucket:          bucket,
		Object:          object,
		ossClient:       ossClient}
}

// PutFromFile 上传文件至阿里云OSS,数据源为本地文件
//
// file 文件路径: /home/upfile.txt 或者 upfile.txt
//
// saveFile 保存文件名,文件名格式,自己可以改,建议保证唯一性
//
// 返回参数 saveURL 成功上传后保存路径
func (o *Client) PutFromFile(file string, saveFile string) (saveURL string, err error) {
	bk, err := o.ossClient.Bucket(o.Bucket)
	if err != nil {
		return
	}
	absPath, err := filepath.Abs(file)
	if err != nil {
		return
	}
	if _, err = os.Stat(absPath); err != nil {
		return
	}
	if saveFile == "" {
		saveFile = filepath.Base(absPath)
	}
	target := filepath.Join(o.Object, saveFile)
	target = filepath.ToSlash(target)
	err = bk.PutObjectFromFile(target, file)
	if err != nil {
		return
	}
	saveURL = fmt.Sprintf("http://%s.%s/%s", o.Bucket, o.EndPoint, target)
	return
}

// PutFromStream 上传文件至阿里云OSS,数据源为本地文件
//
// file 文件流,请求表单提交的文件流
//
// saveFile 保存文件名,文件名格式,自己可以改,建议保证唯一性
//
// 返回参数 saveURL 成功上传后保存路径
func (o *Client) PutFromStream(file *multipart.FileHeader, saveFile string) (saveURL string, err error) {
	bk, err := o.ossClient.Bucket(o.Bucket)
	if err != nil {
		return
	}
	f, e := file.Open()
	if e != nil {
		return "", e
	}
	defer f.Close()

	fd, err := file.Open()
	if err != nil {
		return
	}
	defer fd.Close()
	if saveFile == "" {
		saveFile = filepath.Base(file.Filename)
	}
	target := filepath.Join(o.Object, saveFile)
	target = filepath.ToSlash(target)

	err = bk.PutObject(target, fd)
	if err != nil {
		return
	}
	saveURL = fmt.Sprintf("http://%s.%s/%s", o.Bucket, o.EndPoint, target)
	return
}

// DeleteFile 删除文件
//
// file 文件名称: upfile.txt
//
// 返回参数 err 错误信息
func (o *Client) DeleteFile(file string) (err error) {
	bk, err := o.ossClient.Bucket(o.Bucket)
	if err != nil {
		return
	}
	baseName := filepath.Base(file)
	target := filepath.Join(o.Object, baseName)
	target = filepath.ToSlash(target)
	return bk.DeleteObject(target)
}
