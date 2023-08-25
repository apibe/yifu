package apple

type Argument struct {
	Key             string          `json:"key,required" bson:"key,required"`
	Value           string          `json:"value" bson:"value"`
	ArgumentOpinion ArgumentOpinion `json:"argumentOpinion" json:"argumentOpinion"`
	ArgumentType    ArgumentType    `json:"argumentType" bson:"argumentType"`
	Function        Function        `json:"function" bson:"function"`
	Description     string          `json:"description" bson:"description"`
}

type Arguments []Argument

type Function struct {
	FunctionName string `json:"functionName,required" bson:"functionName,required"`
	Parameter    string `json:"parameter" bson:"parameter"`
}

type ArgumentType string

const (
	Param         ArgumentType = "param"  // url ?拼接
	PathVariables ArgumentType = "path"   // :path 值替换
	Header        ArgumentType = "header" // header
	Body          ArgumentType = "body"   // body
	Auto          ArgumentType = "auto"   // auto 不论类型
)

type ArgumentOpinion string

const (
	Empty    ArgumentOpinion = ""
	Required ArgumentOpinion = "required"
	Static   ArgumentOpinion = "static"
)
