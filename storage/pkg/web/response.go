package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error []string    `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, errors ...string) Response {
	if code >= 300 {
		return Response{code, nil, errors}
	}

	return Response{code, data, nil}
}
