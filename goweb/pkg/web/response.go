package web

type Response struct {
  Code int `json:"code"`
  Data interface{} `json:"data,ommitempty"`
  Error string `json:"error,ommitempty"`
}

func NewResponse(code int, data interface{}, error string) Response {  
  if code < 300 {
    return Response{code, nil, error}
  }
  
  return Response{code, data, ""}
}
