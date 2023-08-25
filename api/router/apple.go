package router

import (
	"github.com/gin-gonic/gin"
	middler "githup.com/apibe/yifu/api/middle"
	"githup.com/apibe/yifu/api/router/apple"
)

func Apple(e *gin.Engine) {
	g := e.Group("/apple")
	{
		g.POST("/upsert/one", middler.ApiMiddleware, apple.UpsertOne)
	}
}
