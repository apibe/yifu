package apple

import (
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/internal/tools/logger"
	"githup.com/apibe/yifu/internal/warehouse/apple"
)

// UpsertOne 根据id更新或删除一条数据
func UpsertOne(c *gin.Context) {
	g := Gin{Ctx: c}
	var p apple.Apple
	err := c.ShouldBindJSON(&p)
	if err != nil {
		g.Error(NotFound)
		return
	}
	logger.Logger.Info(p)
}
