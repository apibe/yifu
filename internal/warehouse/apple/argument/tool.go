package argument

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

func (a *Argument) tool() error {
	server, f, param := a.Function.Parse()
	switch server {
	case "yifu":
		return a.yifu(f, param)
	default:
		return errors.New(fmt.Sprintf("the service %s for %v is not found !", server, a.Key))
	}
}

func (a *Argument) yifu(f string, param url.Values) error {
	switch f {
	case "nowFormat":
		return a.nowFormat(param)
	default:
		return errors.New(fmt.Sprintf("in service yifu, the function %s for %v is not found !", f, a.Key))
	}
}

func (a *Argument) nowFormat(param url.Values) error {
	a.Value = time.Now().Format(param.Get("format"))
	return nil
}
