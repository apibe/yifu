package global

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"githup.com/apibe/yifu/config"
	"githup.com/apibe/yifu/tools/cache"
	"go.uber.org/zap"
)

var (
	Log        *zap.SugaredLogger
	Writer     *rotatelogs.RotateLogs
	C          *config.Server
	CacheItems map[string]cache.Item
	Cache      *cache.Cache
	//MongoDB    *mongo.Client
)
