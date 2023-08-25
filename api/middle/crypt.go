package middler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"githup.com/apibe/yifu/internal/tools/crypt"
	"io/ioutil"
	"net/http"
)

const aesKey = "mRzEAlhF0QTE4KfW"

// CryptMiddleware 加密中间件
func CryptMiddleware(c *gin.Context) {
	// 当 header debug = dreamjoy 时，取消加密步骤直接返回
	_debug := c.Request.Header.Get("debug")
	if _debug == apibe {
		return
	}
	var secretMp map[string]string
	err := c.ShouldBindJSON(&secretMp)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(secretMp[Secret])
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	decrypted, err := crypt.AesDecrypt(decodeBytes, aesKey)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(decrypted))

	c.Next()

	// 加密
	if _debug == apibe {
		return
	}
	msg, _ := c.Get(MSG)
	code, _ := c.Get(CODE)
	result, _ := c.Get(RESULT)
	output := map[string]interface{}{
		CODE:   code,
		MSG:    msg,
		RESULT: result,
	}
	ret, _ := json.Marshal(output)
	origin := ret //变成字节流
	encrypted := crypt.AesEncrypt(origin, aesKey)
	secret := base64.StdEncoding.EncodeToString(encrypted)
	c.Set(Secret, secret)
}
