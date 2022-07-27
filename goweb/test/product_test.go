package test

import (
	"bytes"
	"encoding/json"
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"goweb/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.InitStore("products.json")
	repo := products.InitRepository(db)
	service := products.InitService(repo)
	p := handler.InitProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.CreateProduct())
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

func Test_GetAllProducts_OK(t *testing.T) {

	type producto struct {
		Id            int     `json:"id"`
		Nombre        string  `json:"nombre"`
		Color         string  `json:"color"`
		Precio        float64 `json:"precio"`
		Stock         int     `json:"stock"`
		Codigo        string  `json:"código"`
		Publicado     bool    `json:"publicado"`
		FechaCreacion string  `json:"fecha_de_creación"`
	}
	// Crear el server y definir las rutas
	r := createServer()
	// Crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/products/", "")
	var objRes []producto
	// Indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}

func Test_createProduct_OK(t *testing.T) {

	r := createServer()

	req, rr := createRequestTest(http.MethodPost, "/products/", `{
		"nombre":"Maracuya", "color":"Amarillo", "precio":1233142, "stock":30, "código":"asde312", "publicado":true,
		"fecha_de_creación":"22/06/2022" 
	}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}

func Test_updateProduct_OK(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/products/1", `{
		"nombre":"Guayaba", "color":"Amarillo", "precio":123142, "stock":20, "código":"asde12", "publicado":false,
		"fecha_de_creación":"22/06/2022" 
	}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}

func Test_deleteProduct_OK(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/products/3", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)

}
