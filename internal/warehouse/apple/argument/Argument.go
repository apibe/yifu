package argument

import (
	"fmt"
	"net/url"
	"regexp"
)

type Argument struct {
	Key         string   `json:"key,required" `
	Value       string   `json:"value"`
	Opinion     Opinion  `json:"opinion"`
	Type        Type     `json:"type"`
	Function    Function `json:"function"`
	Description string   `json:"description"`
}

type Opinion string

const (
	Empty     Opinion = ""
	Omitempty Opinion = "omitempty"
	Required  Opinion = "required"
	Static    Opinion = "static"
)

type Type string

const (
	Param         Type = "param"  // url ?拼接
	PathVariables Type = "path"   // :path 值替换
	Header        Type = "header" // header
	Body          Type = "body"   // body
	Auto          Type = "auto"   // auto 不论类型
)

type Function struct {
	Function string `json:"functionName,required"`
}

type Arguments []Argument

func (a *Argument) Assemble() {
	err := a.tool()
	if err != nil {
		fmt.Println(err.Error()) // TODO 处理err
	}
}

// GetValuesEncode 将所有的参数以url参数拼接的形式拼接
func (a *Arguments) GetValuesEncode(argumentType Type) string {
	values := url.Values{}
	for _, arg := range *a {
		if arg.Type == argumentType || arg.Type == Auto {
			values.Add(arg.Key, arg.Value)
		}
	}
	return values.Encode()
}

// Parse yifu@newFormat?aaa=2&bbb=3
func (f *Function) Parse() (server string, function string, param url.Values) {
	rex, _ := regexp.Compile("(\\w+)@(\\w+)\\?(.*)")
	match := rex.FindStringSubmatch(f.Function)
	query, _ := url.ParseQuery(match[2])
	return match[0], match[1], query
}
