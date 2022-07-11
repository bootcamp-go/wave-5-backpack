package web

import (
	"strconv"
)

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	//Si no hay error envío data y el campo de error queda vacío
	if code < 300 {
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	//Si hay error no envío data y envío info del error
	return Response{strconv.FormatInt(int64(code), 10), nil, err}
}
