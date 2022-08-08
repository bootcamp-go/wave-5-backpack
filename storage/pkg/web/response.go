package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"err,omitempty"`
}

func NewResponse(code int, data interface{}, err string) (int, Response) {
	return code, Response{
		Code:  code,
		Data:  data,
		Error: err,
	}
}
