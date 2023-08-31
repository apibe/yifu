package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Init(path string) {
	// todo 配置初始化
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(bytes, &C)
	if err != nil {
		panic(err.Error())
	}
}

type config struct {
	Mod       string   `json:"mod"`
	Addr      []string `json:"addr"`
	MiddleMan []struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
	} `json:"middleMan"`
	Resource resource `json:"resource"`
	Log      string   `json:"log"`
}

var C *config

type resource string

// AppleDirectoryPath directory.json for apple
func (r resource) AppleDirectoryPath() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "bucket", "directory")
}

// ApplePath apple.json for apple
func (r resource) ApplePath(id string) string {
	return fmt.Sprintf("%s/%s", r.AppleDir(), id)
}

// AppleDir apple data dir
func (r resource) AppleDir() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "bucket", "data")
}

// CacheDirectoryPath cache directory path
func (r resource) CacheDirectoryPath() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "cache", "directory")
}

// CacheDir cache data dir
func (r resource) CacheDir() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "cache", "data")
}

// CachePath cache path
func (r resource) CachePath(id string) string {
	return fmt.Sprintf("%s/%s", r.CacheDir(), id)
}
