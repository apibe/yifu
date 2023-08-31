package router

import (
	"github.com/gin-gonic/gin"
)

func Apple(e *gin.Engine) {
	g := e.Group("/apple")
	{
		g.POST("/transport/:name")
	}
}
