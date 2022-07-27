package test

import (
	"bytes"
	"encoding/json"
	"goweb/productos_capas/cmd/server/handler"
	"goweb/productos_capas/internal/products"
	"goweb/productos_capas/pkg/store"
	"log"
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
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/productos")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PATCH("/:id", p.UpdateNamePrice())
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

	// crear el server y definir las rutas
	r := createServer("../products.json")
	// crear request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/productos/", "")

	var objRes Response

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes.Data.([]interface{})) > 0)
}

func TestSaveProduct(t *testing.T) {
	// crear el server y definir las rutas
	r := createServer("../products.json")
	// crear request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/productos/", `{
		"nombre":"Test post","color":"Blanco","precio":100,"stock":55,"codigo":"T001","publicado":true,"fecha_creacion":"13-01-2022"}`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateNamePriceProduct(t *testing.T) {
	r := createServer("../products.json")

	req, rr := createRequestTest(http.MethodPatch, "/productos/7", `{"nombre": "Test patch", "precio": 200}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateProduct(t *testing.T) {
	r := createServer("../products.json")

	req, rr := createRequestTest(http.MethodPut, "/productos/7", `{
		"nombre":"Test update","color":"Azul","precio":300,"stock":20,"codigo":"T002","publicado":false,"fecha_creacion":"04-05-2010"}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDeleteProduct(t *testing.T) {
	r := createServer("../products.json")

	req, rr := createRequestTest(http.MethodDelete, "/productos/7", "")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
