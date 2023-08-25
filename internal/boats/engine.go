package boats

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	Url    string
	Header map[string]string
	Method string
	Body   []byte
	Ctx    context.Context
}

type Response struct {
	Status int
	Header http.Header
	Body   []byte
}

func (req *Request) SentHttp(timeOut time.Duration) (*Response, error) {
	payload := bytes.NewReader(req.Body)
	client := &http.Client{
		Timeout: timeOut,
	}
	if req.Ctx == nil {
		req.Ctx = context.Background()
	}
	request, err := http.NewRequestWithContext(req.Ctx, req.Method, req.Url, payload)
	if err != nil {
		return &Response{http.StatusBadRequest, nil, nil}, err
	}
	for k, v := range req.Header {
		request.Header.Add(k, v)
	}
	resp, err := client.Do(request)
	if err != nil {
		if resp != nil {
			return &Response{resp.StatusCode, resp.Header, nil}, err
		} else {
			return &Response{http.StatusBadGateway, nil, nil}, err
		}
	}
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{resp.StatusCode, resp.Header, nil}, err
	}
	return &Response{resp.StatusCode, resp.Header, body}, err
}
