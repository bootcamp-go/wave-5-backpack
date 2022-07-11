package web

type response struct {
	Code int         `json:"code,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Err  string      `json:"error,omitempty"`
}

func NewRespose(code int, data interface{}, err string) response {
	return response{
		Code: code,
		Data: data,
		Err:  err,
	}
}
