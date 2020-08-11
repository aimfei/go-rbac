package common

const (
	SuccessCode    = "00000"
	SuccessMessage = "success"
)

type Result struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NoDataSuccess() *Result {
	return &Result{Code: SuccessCode, Message: SuccessMessage}
}

func DataSuccess(data interface{}) *Result {
	return &Result{Code: SuccessCode, Message: SuccessMessage, Data: data}
}
