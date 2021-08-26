package common

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    interface{}
}

func NewResponse(success bool, msg string) *Response {
	res := Response{
		success,
		msg,
		nil,
	}
	return &res
}
