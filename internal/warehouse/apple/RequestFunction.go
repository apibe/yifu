package apple

type RequestFunction struct {
	FunctionName string      `json:"functionName,required"`
	Parameter    interface{} `json:"parameter"`
}
