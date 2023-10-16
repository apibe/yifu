package config

type Server struct {
	System   System     `json:"system"`
	Dangguan []Dangguan `json:"dangguan"`
	Resource resource   `mapstructure:"resource" json:"resource" yaml:"resource"`

	// 跨域配置
	Cors CORS   `mapstructure:"cors" json:"cors" yaml:"cors"`
	Log  string `json:"log"`
	// go-cache 配置
	Cache Cache `mapstructure:"cache" json:"cache" yaml:"cache"`

	// redis 配置
	Redis `json:"redis" yaml:"redis"`
	// mongo 配置
	Mongo `json:"mongo" yaml:"mongo"`
	// 远程过程调用 rpc
	Rpc `json:"rpc" yaml:"rpc"`
}
