package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(cod int, data interface{}, err string) Response {
	if cod < 300 {
		return Response{cod, data, ""}
	}
	return Response{cod, nil, err}
}
