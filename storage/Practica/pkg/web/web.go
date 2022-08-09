package web

type Response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
}

func NewRespose(code int, data interface{}, err string) Response {
	return Response{
		Code: code,
		Data: data,
		Err:  err,
	}
}
