package apple

import "github.com/gin-gonic/gin"

func (a *Apple) AutoAssembly(ctx *gin.Context) (body []byte, status int, isCache bool) {
	body = make([]byte, 0)
	status = 400
	isCache = false
	return
}
func (a *Apple) assemblyUrl()      {}
func (a *Apple) assemblyPayload()  {}
func (a *Apple) assemblyArgument() {}
func (a *Apple) assembleArgument() {}
