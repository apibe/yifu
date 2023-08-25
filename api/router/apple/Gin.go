package apple

import (
	"github.com/gin-gonic/gin"
	middler "githup.com/apibe/yifu/api/middle"
	"net/http"
)

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, msg string, data interface{}) {
	g.Ctx.Set(middler.CODE, code)
	g.Ctx.Set(middler.MSG, msg)
	g.Ctx.Set(middler.RESULT, data)
	return
}

func (g *Gin) Success(data interface{}) {
	g.Ctx.Set(middler.CODE, http.StatusOK)
	g.Ctx.Set(middler.MSG, "success")
	g.Ctx.Set(middler.RESULT, data)
}

func (g *Gin) Error(err *Error) {
	details := err.Details()
	g.Ctx.Set(middler.CODE, err.Code())
	g.Ctx.Set(middler.MSG, err.Msg())
	if len(details) > 0 {
		g.Ctx.Set(middler.DETAIL, details)
	}
}
