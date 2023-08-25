package middler

import (
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/internal/tools/crypt"
	"strings"
)

// JWTMiddleware 中间件，检查token
func JWTMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(token, "APIBE t-") {
		token = strings.Replace(token, "APIBE t-", "", 1)
		claims, err := crypt.NewJWT().ParseToken(token) // parseToken 解析token包含的信息
		if err != nil {
			c.Set(MSG, err.Error())
			c.Abort()
			return
		}
		// 设置claims鉴权信息
		c.Set(CLAIMS, claims) // 继续交由下一个路由处理,并将解析出的信息传递下去
	} else {
		c.Set(CODE, ErrReq)
		c.Set(MSG, "token is illegal")
		c.Abort()
		return
	}
	c.Next()
}
