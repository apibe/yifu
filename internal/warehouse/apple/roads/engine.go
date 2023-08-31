package roads

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"
)

// Request
// 请求体
type Request struct {
	Url    string            `json:"url" bson:"url"`
	Header map[string]string `json:"header" bson:"header"`
	Method string            `json:"method" bson:"method"`
	Body   []byte            `json:"body" bson:"body"`
	Ctx    context.Context   `json:"ctx" bson:"-"`
}

// Response
// 返回体
type Response struct {
	Status int         `json:"status" bson:"status"`
	Header http.Header `json:"header" bson:"header"`
	Body   []byte      `json:"body" bson:"body"`
}

// SentHttp 向三方发送请求
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
