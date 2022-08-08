package web

import "github.com/gin-gonic/gin"

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func Response(c *gin.Context, status int, err string, data interface{}) {
	c.JSON(status, response{Data: data, Error: err})
}
