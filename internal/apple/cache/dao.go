package cache

import (
	"encoding/json"
	"githup.com/apibe/yifu/global"
	"os"
)

func Init() {
	err := resetDirectory()
	if err != nil {
		panic(err)
	}
}

type Cache struct {
	ID      string      `json:"id"`
	Timeout int64       `json:"timeout"`
	Claim   interface{} `json:"claim"`
	Base64  string      `json:"base64"`
}

var directory directoryConfig

type directoryConfig map[string]cacheConfig

type cacheConfig struct {
	ID      string      `json:"id"` // id
	Timeout int64       `json:"timeout"`
	Claim   interface{} `json:"claim"` // 载荷
}

func FindById(id string) (*Cache, error) { return read(id) }

func AddOne(cache *Cache) error { return write(cache) }

func RemoveOne(id string) error { return remove(id) }

func read(id string) (cache *Cache, err error) {
	bytes, err := os.ReadFile(global.C.Resource.CachePath(id))
	err = json.Unmarshal(bytes, &cache)
	return
}

func write(cache *Cache) (err error) {
	bytes, err := json.Marshal(cache)
	err = os.WriteFile(global.C.Resource.CachePath(cache.ID), bytes, 0666)
	err = writeCacheToDirectory(&cacheConfig{
		ID:      cache.ID,
		Timeout: cache.Timeout,
		Claim:   cache.Claim,
	})
	return
}

func remove(id string) (err error) {
	err = removeCacheFromDirectory(id)
	err = os.Remove(global.C.Resource.CachePath(id))
	return
}

func readDirectory() (err error) {
	bytes, err := os.ReadFile(global.C.Resource.CacheDirectoryPath())
	err = json.Unmarshal(bytes, &directory)
	if err != nil {
		return
	}
	return
}

func writeDirectory() (err error) {
	bytes, err := json.Marshal(directory)
	if err != nil {
		return
	}
	err = os.WriteFile(global.C.Resource.CacheDirectoryPath(), bytes, 0666)
	if err != nil {
		return
	}
	return readDirectory()
}

func writeCacheToDirectory(cache *cacheConfig) (err error) {
	directory[cache.ID] = *cache
	return writeDirectory()
}

func removeCacheFromDirectory(id string) (err error) {
	delete(directory, id)
	return writeDirectory()
}

func resetDirectory() (err error) {
	err = readDirectory()
	// 判断某个元素是否在某个切片里
	exist := func(flag string, args []string) bool {
		for _, arg := range args {
			if flag == arg {
				return true
			}
		}
		return false
	}
	// 1. 读取 directory 获取所有的 id 和 bucket
	buckets := make([]string, 0)
	for _, c := range directory {
		buckets = append(buckets, c.ID)
	}
	dir, err := os.ReadDir(global.C.Resource.CacheDir())
	for _, entry := range dir {
		if !entry.IsDir() && !exist(entry.Name(), buckets) {
			cache, _ := read(entry.Name())
			directory[entry.Name()] = cacheConfig{
				ID:      cache.ID,
				Timeout: cache.Timeout,
				Claim:   cache.Claim,
			}
		}
	}
	return writeDirectory()
}
