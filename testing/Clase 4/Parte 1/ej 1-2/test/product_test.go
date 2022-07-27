package test

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/cmd/sever/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/file"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("token", "123456")
	fileDB := file.NewFile("../resources/products.json")
	if err := fileDB.Ping(); err != nil {
		log.Fatal(err)
	}
	repository := products.NewRepository(fileDB)
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	router := gin.Default()

	productos := router.Group("/products")
	{
		productos.PUT("/:id", p.UpdateTotal())
		productos.DELETE("/:id", p.Delete())
	}

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("", "")
	return req, httptest.NewRecorder()
}

func TestUpdateProductFuncional(t *testing.T) {
	r := createServer()
	req, res := createRequestTest(http.MethodPut, "/products/2", `{
		"id": 2,
		"name": "Laptop mod",
		"color": "black",
		"price": 999.99,
		"stock": 100,
		"code": "SJD23RFG",
		"published": false,
		"created_at": "2022-06-30"
	}`)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Result().StatusCode)
}

func TestDeleteProductFuncional(t *testing.T) {
	r := createServer()
	req, res := createRequestTest(http.MethodDelete, "/products/7", ``)
	r.ServeHTTP(res, req)
	fmt.Println(res.Body)
	assert.Equal(t, http.StatusNoContent, res.Result().StatusCode)
}
