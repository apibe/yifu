package apple

import (
	"fmt"
	"net/http"
)

type Error struct {
	// 错误码
	code int
	// 错误消息
	msg string
	// 详细信息
	details string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() string {
	return e.details
}

func (e *Error) WithDetails(details string) *Error {
	newError := *e
	newError.details = details

	return &newError
}
func (e *Error) WithMsg(msg string) *Error {
	newError := *e
	newError.msg = msg

	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case NotFound.Code():
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

var (
	Success       = NewError(200, "成功")
	ServerError   = NewError(500, "服务内部错误")
	InvalidParams = NewError(400, "入参错误")

	UnauthorizedAuthNotExist  = NewError(401, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(401, "鉴权失败，Token错误")
	UnauthorizedTokenGenerate = NewError(401, "鉴权失败，Token生成失败")
	UnauthorizedTokenTimeout  = NewError(401, "鉴权失败，Token超时")

	NotFound        = NewError(404, "找不到")
	TooManyRequests = NewError(429, "请求过多")
)
