package config

import "time"

type Cache struct {
	DefaultExpiration time.Duration `json:"default-expiration,omitempty" yaml:"default-expiration"` // 缓存存放时间
	CleanupInterval   time.Duration `json:"cleanup-interval,omitempty" yaml:"cleanup-interval"`     // 缓存清理间隔时间
}
