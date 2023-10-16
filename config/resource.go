package config

import "fmt"

type resource string

// AppleDirectoryPath directory.json for apple
func (r resource) AppleDirectoryPath() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "bucket", "directory")
}

// ApplePath apple.json for apple
func (r resource) ApplePath(id string) string {
	return fmt.Sprintf("%s/%s", r.AppleDir(), id)
}

// AppleDir apple data dir
func (r resource) AppleDir() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "bucket", "data")
}

// CacheDirectoryPath cache directory path
func (r resource) CacheDirectoryPath() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "cache", "directory")
}

// CacheDir cache data dir
func (r resource) CacheDir() string {
	return fmt.Sprintf("%s/%s/%s/%s", r, "apple", "cache", "data")
}

// CachePath cache path
func (r resource) CachePath(id string) string {
	return fmt.Sprintf("%s/%s", r.CacheDir(), id)
}
