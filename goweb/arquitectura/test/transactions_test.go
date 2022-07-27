package test

import (
	"arquitectura/cmd/server/handler"
	"arquitectura/internal/transactions"
	"arquitectura/pkg/store"
	"arquitectura/pkg/web"
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type responseStruct struct {
	Id          int     `json:"id"`
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Reciever    string  `json:"reciever" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "12345")
	db := store.NewStore(pathDB)
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	router := gin.Default()

	tr := router.Group("/transactions")
	tr.Use(TokenAuthMiddleWare())
	tr.POST("/", transactions.Store())
	tr.GET("/", transactions.GetAll())
	tr.PUT("/:id", transactions.Update())
	tr.DELETE("/:id", transactions.Delete())
	tr.PATCH("/:id", transactions.UpdateFields())
	return router
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")
	return req, httptest.NewRecorder()
}

func TokenAuthMiddleWare() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no se encontro el token en variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "missing token"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		c.Next()
	}
}

func TestUpdate(t *testing.T) {
	r := createServer("transactions.json")
	req, rr := createRequest(http.MethodPut, "/transactions/2", `{
		"tranCode": "actualizado2",
		"currency": "CLP",
		"amount": 140.2,
		"transmitter": "carlos",
		"reciever": "catalina",
		"tranDate": "03-09-2021"
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDelete(t *testing.T) {
	r := createServer("transactions.json")
	req, rr := createRequest(http.MethodDelete, "/transactions/3", ``)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
