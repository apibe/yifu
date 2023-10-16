package apple

import (
	"fmt"
	"github.com/google/uuid"
	"githup.com/apibe/yifu/config"
	"testing"
)

func TestFindById(t *testing.T) {
	config.Init("/Users/apple/Desktop/diit/workspace/apibe/yifu/yifu.json")
	apple, err := FindById("4b7236d7-b979-53b4-a7a5-eb12104f3e17")
	if err != nil {
		fmt.Println(apple.ID)
	}
}

func TestFindList(t *testing.T) {
	config.Init("/Users/apple/Desktop/diit/workspace/apibe/yifu/yifu.json")
	list := FindList("bucket")
	fmt.Println(list)
}

func TestAddOne(t *testing.T) {
	config.Init("/Users/apple/Desktop/diit/workspace/apibe/yifu/yifu.json")
	id := uuid.New().String()
	Init()
	AddOne(&Apple{
		ID:          id,
		Bucket:      "bucket",
		Name:        "测试名称",
		Description: "测试描述",
		Method:      "GET",
		Url:         "https://www.baidu.com",
		ContentType: "application/json",
		Payload:     `{"hello":"world"}"`,
		Argument:    nil,
		Format:      nil,
		Status:      0,
	})
}

func TestRemoveOne(t *testing.T) {
	config.Init("/Users/apple/Desktop/diit/workspace/apibe/yifu/yifu.json")
	readDirectory()
	fmt.Println(directory)
	RemoveOne("d58eadeb-ff98-4665-83de-be3fd7d5345d")
}

func TestResetDirectory(t *testing.T) {
	config.Init("/Users/apple/Desktop/diit/workspace/apibe/yifu/yifu.json")
	readDirectory()
}
