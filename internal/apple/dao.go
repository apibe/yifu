package apple

import (
	"encoding/json"
	"githup.com/apibe/yifu/global"
	"os"
)

var directory directoryConfig

type directoryConfig map[string][]appleConfig

type appleConfig struct {
	ID    string      `json:"id"`    // id
	Name  string      `json:"name"`  // 名称
	Claim interface{} `json:"claim"` // 载荷
}

func FindById(id string) (a *Apple, err error) { return read(id) }

func FindList(bucket string) []Apple {
	apples := make([]Apple, 0)
	for _, ad := range directory[bucket] {
		apple, _ := read(ad.ID)
		apples = append(apples, *apple)
	}
	return apples
}

func AddOne(apple *Apple) error { return write(apple) }

func RemoveOne(id string) error { return remove(id) }

func read(id string) (apple *Apple, err error) {
	file, err := os.ReadFile(global.C.Resource.ApplePath(id))
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &apple)
	if err != nil {
		return
	}
	return
}

func write(apple *Apple) (err error) {
	bytes, err := json.Marshal(*apple)
	if err != nil {
		return
	}
	err = os.WriteFile(global.C.Resource.ApplePath(apple.ID), bytes, 0666)
	if err != nil {
		return
	}
	err = writeAppleToDirectory(apple)
	if err != nil {
		return
	}
	return nil
}

func remove(id string) (err error) {
	err = os.Remove(global.C.Resource.ApplePath(id))
	if err != nil {
		return
	}
	err = removeAppleFromDirectory(id)
	if err != nil {
		return
	}
	return err
}

// readDirectory 将 directory.json 写入全局变量 directory
func readDirectory() error {
	global.C.Resource.AppleDir()
	file, err := os.ReadFile(global.C.Resource.AppleDirectoryPath())
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &directory)
}

// writeDirectory 把全局变量 directory 写入 directory.json
func writeDirectory() (err error) {
	bytes, err := json.Marshal(directory)
	if err != nil {
		return
	}
	err = os.WriteFile(global.C.Resource.AppleDirectoryPath(), bytes, 0666)
	if err != nil {
		return
	}
	return readDirectory()
}

// writeAppleToDirectory 把 apple 写入 directory.json 和 directory
func writeAppleToDirectory(apple *Apple) (err error) {
	_ = removeAppleFromDirectory(apple.ID) // 先从索引中删除
	directory[apple.Bucket] = append(
		directory[apple.Bucket], appleConfig{
			ID:    apple.ID,
			Name:  apple.Name,
			Claim: apple.Claim,
		},
	)
	return writeDirectory() // 重新写入索引
}

// removeAppleFromDirectory 从 directory.json 和 directory 删除 apple
func removeAppleFromDirectory(id string) (err error) {
	for k, configs := range directory {
		for i, c := range configs {
			if c.ID == id {
				directory[k] = append(directory[k][:i], directory[k][i+1:]...)
			}
		}
	}
	return writeDirectory()
}

// ResetDirectory 从 bucket/data 重建索引
func ResetDirectory() (err error) {
	err = readDirectory()
	if err != nil {
		return
	}
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
	for _, configs := range directory {
		for _, c := range configs {
			buckets = append(buckets, c.ID)
		}
	}
	// 1. 读取文件夹下所有的文件名
	dir, err := os.ReadDir(global.C.Resource.AppleDir())
	for _, entry := range dir {
		if !entry.IsDir() && !exist(entry.Name(), buckets) {
			apple, _ := read(entry.Name())
			directory[apple.Bucket] = append(directory[apple.Bucket], appleConfig{
				ID:    apple.ID,
				Name:  apple.Name,
				Claim: apple.Claim,
			})
		}
	}
	return writeDirectory()
}
