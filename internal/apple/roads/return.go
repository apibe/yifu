package roads

import (
	"githup.com/apibe/yifu/internal/apple/argument"
)

// todo do grpc

type ResponseFunction struct {
	Name string `json:"name,required"`
}

func (res *Response) Return(resFunc ResponseFunction, arg *argument.Arguments) {
	// todo 加入grpc
	switch resFunc.Name {
	case "dragData":
	case "dragDeepData":
	case "stringToMap":
	case "regexp":
	default:
	}
}
