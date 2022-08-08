package web

//	STRUCT : Response
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

//	FUNCTIONS : NewResponse
func NewResponse(code int, data interface{}, err string) Response {
	if code < 300 {
		return Response{code, data, ""}
	}
	return Response{code, nil, err}

}
