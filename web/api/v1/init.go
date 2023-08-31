package v1

import (
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/cmd/config"
	"githup.com/apibe/yifu/web/api/router"
)

func Init() {
	startup()
}

func startup() {
	e := gin.Default()
	gin.SetMode(config.C.Mod)
	all(e) // 接口注册
	err := e.Run(config.C.Addr...)
	if err != nil {
		panic(err.Error())
	}
}

func all(e *gin.Engine) {
	e.Group("v1")
	router.Apple(e)
	router.Mango(e)
	router.Muskmelon(e)
	router.Orange(e)
	router.Potato(e)
}
