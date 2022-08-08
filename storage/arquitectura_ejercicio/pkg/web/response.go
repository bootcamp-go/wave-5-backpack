package web

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

const (
	ERR_TOKEN_INVALID = "Token inválido"
	ERR_BAD_REQUEST   = "Request inválida, recibe los datos"
	ERR_ID_INVALID    = "El ID que se recibió es inválido, por favor verifique"
	ERR_BAD_INTERNAL  = "Ocurrió un error, lo sentimos :("
)

func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{code, data, ""}
	}
	return Response{code, nil, err}

}
