package common

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

func NewResponse(success bool, msg string) *Response {
	res := Response{
		success,
		msg,
	}
	return &res
}
