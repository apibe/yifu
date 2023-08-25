package apple

type Formats []Format

type Format struct {
	Condition string `json:"condition"` // $.code!=0  || $.code!exist
	Template  string `json:"template"`  // 格式化预设内容
	Cache     Cache  `json:"cache"`     // 缓存
}
