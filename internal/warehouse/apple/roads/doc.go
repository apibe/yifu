package roads

var leaves = []struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}{
	{"AES", "对请求体进行AES加密,参数是secretKey"},
	{"MD5", "对请求体进行MD5加密"},
	{"BASE64", "对请求体进行BASE64加密"},
}

// GetLeaves
// 返回所有的 leaves case
func GetLeaves() interface{} {
	return leaves
}

var returns = []struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}{
	{"dragData", "对一层"},
	{"dragDeepData", ""},
	{"stringToMap", ""},
	{"regexp", `支持gjson语法,正则替换：{"$.data.name":{"matching":"","replace"}}`},
}

// GetReturns
// 获取所有的 return case
func GetReturns() interface{} {
	return returns
}
