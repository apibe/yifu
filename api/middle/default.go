package middler

import (
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/internal/tools/logger"
)

const (
	ErrOK    = iota // 无误
	Err             // 未知err
	ErrReq          // 请求内容报错
	ErrRes          // 返回数据序列化错误
	ErrMgo          // mongo 运行错误
	ErrRedis        // redis 运行错误
	ErrMysql        // mysql允许错误
	ErrAuth         // 权限内容报错
)

const (
	CODE   = "code"   // 响应状态描述标识
	MSG    = "msg"    // 错误日志标识
	RESULT = "result" // 正确返回内容标识
	DETAIL = "detail" // 详细错误信息

	CLAIMS = "claims" // 鉴权信息标识
	Secret = "secret" // secret标识
	apibe  = "APIBE"  // debug标识
)

func ApiMiddleware(c *gin.Context) {
	// 设置初始的返回体
	c.Set(CODE, ErrOK)
	c.Set(MSG, "")
	c.Set(RESULT, "")

	c.Next()

	// 响应后
	code, _ := c.Get(CODE)
	msg, _ := c.Get(MSG)
	result, _ := c.Get(RESULT)
	if code != ErrOK {
		logger.Logger.Errorf("request err: code:%v msg:%v", code, msg)
	}

	if secret, _ := c.Get(Secret); secret != nil {
		c.JSON(c.Writer.Status(), gin.H{
			Secret: secret,
		})
		return
	}

	// 只有header里填入了debug的才会以原文显示
	c.JSON(c.Writer.Status(), gin.H{
		CODE:   code,
		MSG:    msg,
		RESULT: result,
	})
}
