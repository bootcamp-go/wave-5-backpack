package web

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
	Code  int         `json:"code"`
}

func NewResponse(data interface{}, err string, code int) *Response {
	if code < 300 {
		return &Response{Data: data, Error: "", Code: code}
	}

	return &Response{Data: nil, Error: err, Code: code}
}
