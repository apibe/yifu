package middler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// White 接口访问白名单
var White = []string{"172.8.1.1"}

func WhiteMiddleware(c *gin.Context) {
	flag := false
	for _, white := range White {
		if c.ClientIP() == white {
			flag = true
		}
	}
	if flag {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusNonAuthoritativeInfo)
	}
}
