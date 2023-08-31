package apple

import (
	"errors"
	"fmt"
	"githup.com/apibe/yifu/internal/warehouse/apple/argument"
	"githup.com/apibe/yifu/internal/warehouse/apple/roads"
	"net/http"
	"strings"
	"time"
)

func (a *Apple) AutoAssembly(catch map[string]string, timeOut time.Duration) (body []byte, status int, isCache bool, err error) {
	req := &roads.Request{}
	req.Header = make(map[string]string) // 初始化map
	err = a.valueReplace(catch)
	if err != nil {
		return nil, 0, false, err
	}
	// 1. 组装 Apple 的 argument
	a.assemblyArgument()
	// 2. 尝试读取缓存
	a.assemblyCache()
	if bytes, err := a.Cache.CacheGet(); err == nil {
		return bytes, http.StatusOK, true, err
	}
	// 3. 拼装url
	a.assemblyUrl(req)
	// 4. 拼装 Content-Type
	a.assemblyContentType(req)
	// 5. 拼装 header
	a.assemblyHeader(req)
	// 6. 拼装请求体 payload
	a.assemblyPayload(req)
	// 7. 处理请求
	a.doRequestFunction(req)
	// 8. 执行请求
	res, err := req.SentHttp(timeOut)
	if err != nil {
		if res != nil {
			return nil, res.Status, false, err
		} else {
			return nil, http.StatusBadGateway, false, err
		}
	}
	// 9. 处理响应体
	a.doResponseFunction(res)
	// 10. 格式化相应内容，并进行缓存
	a.doFormat(res)
	return nil, 200, true, nil
}

// 把请求体内容和Apple相结合
func (a *Apple) valueReplace(catch map[string]string) error {
	for key, value := range catch {
		for i, arg := range a.Argument {
			if key == arg.Key && arg.Opinion != argument.Static {
				a.Argument[i].Value = value
			}
			if catch[arg.Key] == "" && arg.Opinion == argument.Required {
				return errors.New(fmt.Sprintf("key %v is required", arg.Key))
			}
		}
	}
	return nil
}

func (a *Apple) assemblyCache() {
	id := a.Name + "-" + a.Cache.ID
	for _, arg := range a.Argument {
		id = strings.Replace(id, "$."+arg.Key, arg.Value, -1)
	}
	a.Cache.ID = id
}

// 1. 拼装 argument
func (a *Apple) assemblyArgument() {
	for i, arg := range a.Argument {
		if arg.Function.Name != "" {
			a.Argument[i].Assemble()
		}
	}
}

// 2. 拼装 url
func (a *Apple) assemblyUrl(req *roads.Request) {
	uri := a.Url
	// 1. 拼接需要添加的 url 后半部分
	param := a.Argument.GetValuesEncode(argument.Param)
	// 2. 组装 url param
	if param != "" {
		if !strings.Contains(a.Url, "?") {
			uri = fmt.Sprint(a.Url, "?", param)
		} else {
			uri = fmt.Sprint(a.Url, param)
		}
	}
	// 3. 组装 url pathValue  http://www.baidu.com/:aa/:dd/ss
	for _, arg := range a.Argument {
		uri = strings.Replace(uri, fmt.Sprintf(":%v", arg.Key), arg.Value, 1)
	}
	a.Url = uri
	req.Url = uri
}

// 3. 拼装 method
func (a *Apple) assemblyMethod(req *roads.Request) {
	if a.Method != "" {
		req.Method = a.Method
	}
}

// 4. 拼装 Content-Type
func (a *Apple) assemblyContentType(req *roads.Request) {
	if a.ContentType != "" {
		req.Header["Content-Type"] = a.ContentType
	}
}

// 5. 拼装 header
func (a *Apple) assemblyHeader(req *roads.Request) {
	reqHeader := make(map[string]string)
	arguments := a.Argument
	for _, arg := range arguments {
		if arg.Type == argument.Header || arg.Type == argument.Auto {
			reqHeader[arg.Key] = arg.Value
		}
	}
	req.Header = reqHeader
}

// 6.组装请求体 payload
func (a *Apple) assemblyPayload(req *roads.Request) {
	const (
		None             = ""
		ApplicationXml   = "application/xml"
		ApplicationJson  = "application/json"
		ApplicationJS    = "application/javascript"
		ApplicationXForm = "application/x-www-form-urlencoded"
		FormData         = "multipart/form-data"
		Text             = "text/plain"
		Html             = "text/html"
	)

	switch a.ContentType {
	case None:
		reqBody := []byte(a.Argument.GetValuesEncode(argument.Body))
		a.Payload = string(reqBody)
		req.Body = reqBody
		break
	case ApplicationJson:
		str := a.Payload
		for _, arg := range a.Argument {
			value := strings.TrimSpace(arg.Value)
			if arg.Type == argument.Body || arg.Type == argument.Auto {
				str = strings.Replace(str, "$."+arg.Key, value, -1)
			}
		}
		//fmt.println(str)
		a.Payload = str
		req.Body = []byte(str)
		break
	case ApplicationXForm:
		reqBody := []byte(a.Argument.GetValuesEncode(argument.Body))
		a.Payload = string(reqBody)
		req.Body = reqBody
		break
	case ApplicationXml:
		str := a.Payload
		for _, arg := range a.Argument {
			if arg.Type == argument.Body {
				str = strings.Replace(str, "$."+arg.Key, arg.Value, -1)
			}
		}
		a.Payload = str
		req.Body = []byte(str)
		break
	default:
		reqBody := []byte(a.Argument.GetValuesEncode(a.Payload))
		a.Payload = string(reqBody)
		req.Body = reqBody
		break
	}
}

// 7. 执行 请求前 function
func (a *Apple) doRequestFunction(req *roads.Request) {
	if a.RequestFunction.Function != "" {
		req.Leave(a.RequestFunction, &a.Argument)
	}
}

// 8. 执行 请求后 function
func (a *Apple) doResponseFunction(res *roads.Response) {
	if a.ResponseFunction.Name != "" {
		res.Return(a.ResponseFunction, &a.Argument)
	}
}

// 9. 执行 最终格式化
func (a *Apple) doFormat(res *roads.Response) {
	if res.Status != http.StatusOK {
		return
	} else {
		res.Body = a.Format.Format(a.Cache, res.Body)
	}
}
