package apple

import (
	"fmt"
	"githup.com/apibe/yifu/internal/tools/config"
)

var All map[string]map[string]Apple // bucket.uuid.apple

func Init() {
	resource := config.C.Resource
	fmt.Println(resource)
	// todo apples 初始化
}

func FindById(id string) *Apple                           { return nil }
func FindByGroupAndName(group string, name string) *Apple { return nil }
func FindList(group string, name string) []Apple          { return nil }

func AddOne(apple *Apple) string                     { return "" }
func UpdateStatusById(id string, status Status) bool { return false }
func DelById() bool                                  { return false }
