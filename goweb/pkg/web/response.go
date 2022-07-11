package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, error string) Response {
	if code >= 300 {
		return Response{code, nil, error}
	}

	return Response{code, data, ""}
}
