package router

import "github.com/gin-gonic/gin"

func Mango(e *gin.Engine) {
	e.Group("mango")
}
