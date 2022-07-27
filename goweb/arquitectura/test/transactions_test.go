package test

import (
	"arquitectura/cmd/server/handler"
	"arquitectura/internal/domain"
	"arquitectura/internal/transactions"
	"arquitectura/pkg/store"
	"arquitectura/pkg/web"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
		"tranCode": "actualizado3",
		"currency": "USD",
		"amount": 140.2,
		"transmitter": "carlos",
		"reciever": "catalina",
		"tranDate": "03-09-2021"
	}`)

	r.ServeHTTP(rr, req)

	tes := domain.Transaction{
		Id:          2,
		TranCode:    "actualizado3",
		Currency:    "USD",
		Amount:      140.2,
		Transmitter: "carlos",
		Reciever:    "catalina",
		TranDate:    "03-09-2021",
	}

	expectedResponse, err := json.Marshal(web.NewResponse(200, tes, ""))

	assert.Equal(t, 200, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, rr.Body.Bytes())
}

func TestDelete(t *testing.T) {
	r := createServer("transactions.json")
	req, rr := createRequest(http.MethodDelete, "/transactions/2", ``)

	r.ServeHTTP(rr, req)

	expectedResponse, err := json.Marshal(web.NewResponse(200, fmt.Sprintf("la transaccion con id %d ha sido eliminada", 2), ""))

	assert.Equal(t, 200, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, rr.Body.Bytes())
}
