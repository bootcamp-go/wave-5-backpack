package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := store.NewStore("products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.PATCH("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdateProductOk(t *testing.T) {
	productExpected := map[string]interface{}{"ID": 4.00,
		"Nombre":        "Televisor 45 SAMSUNG",
		"Color":         "Negro",
		"Precio":        90000.00,
		"Stock":         200.00,
		"Codigo":        "AR65RT0",
		"Publicado":     true,
		"FechaCreacion": "2022-07-10 21:21:01.892906 -0300 -03"}

	product := web.Response{}

	r := createServer()

	req, rr := createRequestTest(http.MethodPatch, "/products/4", `{"Nombre": "Televisor 45 SAMSUNG", "Precio": 90000}`)
	//reqE, rrE := createRequestTest(http.MethodPatch, "/products/1", `{"Nombre": "Televisor 45 SAMSUNG", "Precio": "90000"}`)

	r.ServeHTTP(rr, req)
	//r.ServeHTTP(rrE, reqE)

	//assert.Equal(t, 200, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &product)
	assert.Nil(t, err)
	//assert.NotEmpty(t, product)
	assert.Equal(t, productExpected, product.Data)
	//assert.Equal(t, 400, rrE.Code)
}

func TestDeleteProductOk(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodDelete, "/products/2", `{}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
