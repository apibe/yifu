package v1

import (
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/internal/tools/config"
)

func Init() {
	e := gin.Default()
	gin.SetMode(config.C.Mod)
	all(e) // 接口注册
	err := e.Run(config.C.Addr...)
	if err != nil {
		panic(err.Error())
	}
}

func setupRouter(engine *gin.Engine) {

}

func startHttpServer() error {
	engine := gin.Default()
	setupRouter(engine)
	if err := engine.Run(config.C.Addr...); err != nil {
		return err
	}
	return nil
}
