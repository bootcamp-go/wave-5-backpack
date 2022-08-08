package web

type Reponse struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Reponse {
	if code < 300 {
		return Reponse{code, data, ""}
	}
	return Reponse{code, nil, err}
}
