package apple

import "githup.com/apibe/yifu/internal/warehouse/apple/cache"

func Init() {
	err := resetDirectory()
	cache.Init()
	if err != nil {
		panic(err.Error())
	}
}
