package apple

import (
	"githup.com/apibe/yifu/internal/apple/cache"
)

func Init() {
	err := ResetDirectory()
	cache.Init()
	if err != nil {
		panic(err.Error())
	}
}
