package web

const (
	STATUS_OK    = "OK"
	STATUS_ERROR = "ERROR"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func NewResponse(data interface{}, err string) Response {
	var status string
	if data != nil {
		status = STATUS_OK
	} else {
		status = STATUS_ERROR
	}
	return Response{
		Status: status,
		Data:   data,
		Error:  err,
	}
}
