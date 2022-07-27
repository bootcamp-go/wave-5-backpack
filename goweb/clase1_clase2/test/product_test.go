package test

import (
	"bytes"
	"encoding/json"
	"goweb/clase1_clase2/cmd/server/handler"
	"goweb/clase1_clase2/internal/products"
	"goweb/clase1_clase2/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := store.NewStore(pathDB)
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PATCH("/:id", p.UpdateFields())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllProducts(t *testing.T) {

	type Response struct {
		Code  int         `json:"code"`
		Data  interface{} `json:"data,omitempty"`
		Error string      `json:"error,omitempty"`
	}

	// crear el Server y definir las Rutas
	r := createServer("../products.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, res := createRequestTest(http.MethodGet, "/products/", "")

	var objRes Response

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	err := json.Unmarshal(res.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data.([]interface{})) > 0)
}

func TestSaveProduct(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("../products.json")
	// crear Request del tipo POST y Response para obtener el resultado
	req, res := createRequestTest(http.MethodPost, "/products/", `{
        "nombre":"Lapicero","color":"Negro","precio":200,"stock":4,"codigo":"L7865","publicado":true,"fecha":"2022-01-07"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestUpdateNamePriceProduct(t *testing.T) {
	r := createServer("../products.json")

	req, res := createRequestTest(http.MethodPatch, "/products/11",
		`{"nombre": "Lapices", "precio":500}`)

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestUpdateProduct(t *testing.T) {
	r := createServer("../products.json")

	req, res := createRequestTest(http.MethodPut, "/products/11",
		`{"nombre":"Marcadores","color":"Rojo","precio":455,"stock":2,"codigo":"M654","publicado":false,"fecha":"2022-03-07"}`)

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}

func TestDeleteProduct(t *testing.T) {
	r := createServer("../products.json")

	req, res := createRequestTest(http.MethodDelete, "/products/11", "")

	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)
}
