package apple

type Cache struct {
	//接口uuid + 请求内容
	CacheOpen bool `json:"cacheOpen"` // 是否开启缓存
	Hold      int  `json:"hold"`      // 缓存时长 s
}
