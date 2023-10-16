package cache

import "time"

type Config struct {
	ID   string `json:"id"`   // cache id
	Open bool   `json:"open"` // 是否开启缓存
	Hold int64  `json:"hold"` // 缓存时长 s
}

func (c *Config) CacheGet() ([]byte, error) {
	cache, err := read(c.ID)
	if err != nil {
		return nil, err
	}
	return []byte(cache.Base64), err
}

// CacheSet
// 设置 apple 缓存数据
func (c *Config) CacheSet(value []byte, claim interface{}) error {
	return write(&Cache{
		ID:      c.ID,
		Timeout: time.Now().Unix() + c.Hold,
		Claim:   claim,
		Base64:  string(value),
	})
}

// CacheRemoveOne
// 移除 apple 缓存数据
func (c *Config) CacheRemoveOne() error {
	return remove(c.ID)
}

// CacheRemoveAll
// 移除所有的缓存
func (c *Config) CacheRemoveAll() int64 {
	res := int64(0)
	for _, cfg := range directory {
		err := remove(cfg.ID)
		if err == nil {
			res++
		}
	}
	return res
}
