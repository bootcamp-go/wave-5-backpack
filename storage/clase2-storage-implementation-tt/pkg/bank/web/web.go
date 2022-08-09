package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//	STRUCT : Response
type ResponseSt struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

//	FUNCTIONS : NewResponse
func NewResponse(code int, data interface{}, err string) ResponseSt {
	if code < 300 {
		return ResponseSt{code, data, ""}
	}
	return ResponseSt{code, nil, err}

}

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}
