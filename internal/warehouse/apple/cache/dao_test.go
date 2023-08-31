package cache

import (
	"fmt"
	"githup.com/apibe/yifu/cmd/config"
	"testing"
)

const cpath = "/Users/apple/Desktop/diit/workspace/apibe/yifu/yifu.json"

func TestResetDirectory(t *testing.T) {
	config.Init(cpath)
	resetDirectory()
}

func TestRead(t *testing.T) {
	config.Init(cpath)
	resetDirectory()
	cache, err := read("hello-4b7236d7-b979-53b4-a7a5-eb12104f3e17")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cache)
}

func TestFindById(t *testing.T) {
	config.Init(cpath)
	resetDirectory()
	cache, err := FindById("hello-4b7236d7-b979-53b4-a7a5-eb12104f3e17")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(cache)
}

func TestAddOne(t *testing.T) {
	config.Init(cpath)
	resetDirectory()
	AddOne(&Cache{
		ID:      "aaasssddd",
		Timeout: 100000000,
		Claim:   map[string]interface{}{"1": "2"},
		Base64:  "claim",
	})
}

func TestRemove(t *testing.T) {
	config.Init(cpath)
	resetDirectory()
	RemoveOne("aaasssddd")
}

func TestFMT(t *testing.T) {
	fmt.Println(fmt.Sprint("a", "b"))
}
