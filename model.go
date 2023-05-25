package core

import (
	"time"

	"github.com/gin-gonic/gin"
)

func JsonSuccess(data any) *JsonResponse {
	var t int64
	if gin.ReleaseMode != gin.Mode() {
		t = time.Now().Unix()
	} else {
		t = 0
	}
	return &JsonResponse{
		Code: 0,
		Msg:  "",
		Data: data,
		Time: t,
	}
}

func JsonFail(code int, msg string) *JsonResponse {
	var t int64
	if gin.ReleaseMode != gin.Mode() {
		t = time.Now().Unix()
	} else {
		t = 0
	}
	return &JsonResponse{
		Code: code,
		Msg:  msg,
		Data: nil,
		Time: t,
	}
}

type JsonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
	Time int64  `json:"time,omitempty"`
}
