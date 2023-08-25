package v1

import (
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/api/router"
)

func all(e *gin.Engine) {
	e.Group("v1")
	router.Apple(e)
	router.Mango(e)
	router.Muskmelon(e)
	router.Orange(e)
	router.Potato(e)
}
