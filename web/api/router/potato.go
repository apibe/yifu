package router

import "github.com/gin-gonic/gin"

func Potato(e *gin.Engine) {
	e.Group("potato")
}
