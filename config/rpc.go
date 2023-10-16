package config

type Rpc struct {
	RequestFunction  string `json:"before-request" yaml:"before-request"`
	ResponseFunction string `json:"response-function" yaml:"response-function"`
}
