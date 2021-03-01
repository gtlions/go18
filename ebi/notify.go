package ebi

import (
	"net/http"

	"github.com/gorilla/schema"
)

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

// ParseRefuncNotify 解析退款异步通知的参数
//
// req：*http.Request
//
// 返回参数 rsp 请求的参数
//
// 返回参数 err 错误信息
//
func ParseRefuncNotify(req *http.Request) (rsp *NotifyRefundRequest, err error) {
	var decoder = schema.NewDecoder()
	rspTemp := NotifyRefundRequest{}
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
