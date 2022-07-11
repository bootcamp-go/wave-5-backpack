package web

import "fmt"

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	if code == 500 {
		fmt.Printf("Error Interno: %s", err)
		return Response{500, data, "internal server error"}
	}
	if code < 300 {
		return Response{code, data, ""}
	}
	return Response{code, nil, err}
}
