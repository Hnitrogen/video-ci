package app

import "github.com/gin-gonic/gin"

// 封装一下响应
type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode, code int, Msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: code,
		Msg:  Msg,
		Data: data,
	})
	return
}
