package initialize

import (
	"encoding/json"
	"githup.com/apibe/yifu/global"
	"os"
)

func Config(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(bytes, &global.C)
	if err != nil {
		panic(err.Error())
	}
}
