package main

import (
	"flag"
	"server/tools/apibe/factory/internal/tools/config"
	"server/tools/apibe/factory/internal/warehouse/apple"
)

var path = flag.String("conf", "./factory.toml", "string类型参数")

func init() {
	flag.Parse()
	config.Init(*path)
	apple.Init()
}

func main() {}
