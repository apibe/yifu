package middler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	allowOrigin      = "*"
	allowMethod      = "POST, GET, OPTIONS, PUT, DELETE,UPDATE"
	allowHeader      = "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type,debug"
	exposeHeaders    = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Cache,Expires,Last-Modified,Pragma,FooBar"
	maxAge           = "172800"
	AllowCredentials = "false"
)

// CrossMiddleware 跨域访问：cross origin resource share
func CrossMiddleware(context *gin.Context) {
	//method := context.Request.Method
	context.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
	context.Header("Access-Control-Allow-Origin", allowOrigin) // 设置允许访问所有域
	context.Header("Access-Control-Allow-Methods", allowMethod)
	context.Header("Access-Control-Allow-Headers", allowHeader)
	context.Header("Access-Control-Expose-Headers", exposeHeaders)
	context.Header("Access-Control-Max-Age", maxAge)
	context.Header("Access-Control-Allow-Credentials", AllowCredentials)
	context.Set("content-type", "application/json")
	if context.Request.Method == "OPTIONS" {
		context.AbortWithStatus(http.StatusNoContent)
	}
	//处理请求
	context.Next()
}
