package roads

import "githup.com/apibe/yifu/internal/warehouse/apple/argument"

// todo do grpc

type RequestFunction struct {
	Function string `json:"function,required"`
}

func (req *Request) Leave(reqFunc RequestFunction, arg *argument.Arguments) {
	switch reqFunc.Function {
	case "AES":
	case "BASE64":
	case "MD5":
	default:
	}
}
