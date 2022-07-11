package web

// --------------------------------------------
// --------------- Estructuras ----------------
// --------------------------------------------

// Clase 4 Ejercicio 1 Parte 1
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// --------------------------------------------
// ----------------- Métodos ------------------
// --------------------------------------------

func NewResponse(code int, data interface{}, error string) *Response {
	if code < 300 {
		return &Response{
			Code:  code,
			Data:  data,
			Error: "",
		}
	}
	return &Response{
		Code:  code,
		Data:  nil,
		Error: error,
	}
}
