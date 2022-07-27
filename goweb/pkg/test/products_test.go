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

// ----------------------------------------
// Funciones necesarias para realizar los Test
// ----------------------------------------

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123")

	db := store.NewStore("./products_test.json")
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	handler := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	{
		pr.PUT("/:id", handler.Update())
		pr.DELETE("/:id", handler.Delete())
	}

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123")

	return req, httptest.NewRecorder()
}

// ----------------------------------------
// Tests
// ----------------------------------------

func Test_Update_OK(t *testing.T) {
	// Crear el Server y definir las rutas
	r := createServer()
	// Crear Request del tipo PUT y Response para obtener el resultado
	req, resRec := createRequestTest(http.MethodPut, "/products/1", `{
		"nombre": "Producto 100",
		"color": "rojo",
		"precio": 100,
		"stock": 10,
		"codigo": "123",
		"publicado": true,
		"fechaCreacion": "2020-01-01"
	}`)

	// Indicar al servidor que puede atender la petición
	r.ServeHTTP(resRec, req)

	// Verificaciones
	assert.Equal(t, http.StatusOK, resRec.Code)
	var obj interface{}
	err := json.Unmarshal(resRec.Body.Bytes(), &obj)
	assert.Nil(t, err)
}

func Test_Delete_OK(t *testing.T) {
	// Crear el Server y definir las rutas
	r := createServer()
	// Crear Request del tipo DELETE y Response para obtener el resultado
	req, resRec := createRequestTest(http.MethodDelete, "/products/2", "")

	// Indicar al servidor que puede atender la petición
	r.ServeHTTP(resRec, req)

	// Verificaciones
	assert.Equal(t, http.StatusOK, resRec.Code)
}
