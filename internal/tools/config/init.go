package config

func Init(path string) {
	// todo 配置初始化
}

type config struct {
	Mod       string   `json:"mod"`
	Addr      []string `json:"addr"`
	MiddleMan []struct {
		Name string `json:"name"`
		Addr string `json:"addr"`
	} `json:"middleMan"`
	Resource string `json:"resource"`
	Log      string `json:"log"`
}

var C *config
