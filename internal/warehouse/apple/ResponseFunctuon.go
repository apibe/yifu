package apple

type ResponseFunction struct {
	FunctionName string      `json:"functionName,required"`
	Parameter    interface{} `json:"parameter"`
}
