package web

type Response struct {
	Code  int
	Data  interface{}
	Error string
}

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{code, data, ""}
	}

	return Response{code, nil, err}
}
