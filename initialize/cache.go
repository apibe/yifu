package initialize

import (
	"githup.com/apibe/yifu/global"
	"githup.com/apibe/yifu/tools/cache"
)

func Cache() {
	global.Cache = cache.NewFrom(global.C.Cache.DefaultExpiration, global.C.Cache.CleanupInterval, global.CacheItems)
}
