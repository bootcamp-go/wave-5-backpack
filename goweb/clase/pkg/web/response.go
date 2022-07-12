package web

type response struct {
	Code  int         `json:"code"`
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

func NewResponse(code int, data interface{}, err string) response {
	if code >= 300 {
		return response{
			Code:  code,
			Error: err,
		}
	}
	return response{
		Code: code,
		Data: data,
	}
}
