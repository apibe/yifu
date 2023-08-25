package apple

type Apple struct {
	ID               string           `json:"id"`
	Bucket           string           `json:"bucket"`
	Apple            string           `json:"apple"`
	Description      string           `json:"description"`
	Method           string           `json:"method,required"`
	Url              string           `json:"url,required"`
	ContentType      string           `json:"Content-Type"`
	Payload          string           `json:"payload"`
	Argument         Arguments        `json:"argument"`
	RequestFunction  RequestFunction  `json:"requestFunction"`
	ResponseFunction ResponseFunction `json:"responseFunction"`
	Format           Formats          `json:"format"`
	Cache            Cache            `json:"cache"`
	Status           Status           `json:"status"`
}

type Status uint8

const (
	NeedCheck Status = iota // 需要审查
	Working                 // 已被审查 工作中
	Unused                  // 弃用  弃用

	Deleted // 逻辑删除 (不可见) 在程序编写中总应该是最大的一位数字
)
