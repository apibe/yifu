package formats

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"githup.com/apibe/yifu/internal/warehouse/apple/cache"
	"reflect"
	"strconv"
	"strings"
)

type Formats []Format

type Format struct {
	Condition string `json:"condition"` // $.code!=0  || $.code!exist
	Template  string `json:"template"`  // 格式化预设内容
	CacheOpen bool   `json:"cacheOpen"` // 缓存
}

func (fs *Formats) Format(cache cache.Config, body []byte) []byte {
	for _, f := range *fs {
		var fmtJsonMap map[string]interface{}
		err := json.Unmarshal([]byte(f.Template), &fmtJsonMap)
		if err != nil {
			continue
		}
		// 判断条件是否成立
		if isCondition(f.Condition, string(body)) {
			format := responseFmt(fmtJsonMap, string(body))
			marshal, _ := json.Marshal(format)
			if f.CacheOpen {
				_ = cache.CacheSet(marshal, nil)
			}
			return marshal
		}
	}
	return body
}

// 判断条件是否成立
func isCondition(condition string, gson string) bool {
	condition = strings.Replace(condition, " ", "", -1)
	//fmt.println(condition)
	if strings.Contains(condition, "=") {
		if strings.Contains(condition, "!=") {
			split := strings.Split(condition, "!=")
			if strings.HasPrefix(split[0], "$.") {
				trimPrefix := strings.TrimPrefix(split[0], "$.")
				return gjson.Parse(gson).Get(trimPrefix).String() != split[1]
			} else {
				trimPrefix := strings.TrimPrefix(split[1], "$.")
				return gjson.Parse(gson).Get(trimPrefix).String() != split[0]
			}
		} else {
			split := strings.Split(condition, "=")
			if strings.HasPrefix(split[0], "$.") {
				trimPrefix := strings.TrimPrefix(split[0], "$.")
				return gjson.Parse(gson).Get(trimPrefix).String() == split[1]
			} else {
				trimPrefix := strings.TrimPrefix(split[1], "$.")
				return gjson.Parse(gson).Get(trimPrefix).String() == split[0]
			}
		}
	}
	if strings.Contains(condition, "exist") {
		if strings.Contains(condition, "!exist") {
			condition = strings.Replace(condition, "!exist", "", 1)
			trimPrefix := strings.TrimPrefix(condition, "$.")
			return gjson.Parse(gson).Get(trimPrefix).Exists()
		} else {
			condition = strings.Replace(condition, "exist", "", 1)
			trimPrefix := strings.TrimPrefix(condition, "$.")
			return gjson.Parse(gson).Get(trimPrefix).Exists()
		}
	}
	return false
}

//根据fmtResponse格式化返回值  string、float、value
func responseFmt(fmtJsonMap map[string]interface{}, gson string) map[string]interface{} {
	for k, v := range fmtJsonMap {
		//如果值是string,就取值
		if str, ok := v.(string); ok {
			if strings.HasPrefix(str, "$.") {
				i1 := strings.Index(str, ".")
				i2 := strings.LastIndex(str, ".")
				splitN := strings.Split(str, ".")
				value := gjson.Parse(gson).Get(str[i1+1 : i2])
				switch splitN[len(splitN)-1] {
				case "string":
					fmtJsonMap[k] = value.String()
					break
				case "float":
					fmtJsonMap[k] = value.Float()
					break
				case "int":
					fmtJsonMap[k] = value.Int()
					break
				case "bool":
					fmtJsonMap[k] = value.Bool()
					break
				default:
					fmtJsonMap[k] = value.Value()
				}
			}
		}
		//如果值是map就在挖一层
		if mp, ok := v.(map[string]interface{}); ok {
			//对于map直接替换
			ddd := responseFmt(mp, gson)
			fmtJsonMap[k] = ddd
		}
		//如果值是[]就在挖两层,最多只处理一层json5737707
		if slices, ok := v.([]interface{}); ok {
			if mapp1, ok := slices[0].(map[string]interface{}); ok {
				sliceNil := isSlice(mapp1, gson)
				fmtJsonMap[k] = sliceNil
			}
		}
	}
	return fmtJsonMap
}

//如果是切片则执行这部分内容
func isSlice(getMap map[string]interface{}, gson string) []interface{} {
	finalSlice := make([]interface{}, 0)
	nilMap := make(map[string]interface{})
	//创建一个空map，用来停止格式化
	for k := range getMap {
		nilMap[k] = nil
	}
	for i := 0; i < 1000; i++ { // 为避免出现性能问题,最大处理1000条的slice
		newMap := make(map[string]interface{})
		for i2, i3 := range getMap {
			if str, ok := i3.(string); ok && strings.Contains(str, "[#]") {
				newMap[i2] = strings.Replace(str, "[#]", strconv.Itoa(i), -1)
			}
		}
		m := responseFmt(newMap, gson)
		if reflect.DeepEqual(nilMap, m) {
			break
		}
		finalSlice = append(finalSlice, m)
	}
	return finalSlice
}
