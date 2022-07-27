package test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"testing/4/tm/test_funcional/cmd/server/handler"
	"testing/4/tm/test_funcional/internal/repository"
	"testing/4/tm/test_funcional/internal/service"
	"testing/4/tm/test_funcional/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	err := os.Setenv("TOKEN", "123")
	if err != nil {
		log.Fatal(err)
	}

	db := store.NewStore("./products.json")
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s)

	router := gin.Default()

	products := router.Group("/products")
	products.PUT("/:id", h.Update())
	products.DELETE("/:id", h.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123")

	return req, httptest.NewRecorder()
}

func TestUpdateOK(t *testing.T) {
	r := createServer()

	// requiere crear primero el elemento en products.json
	req, res := createRequestTest(http.MethodPut, "/products/1", `{"name":"banana", "price":1.25, "quantity":10}`)

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteOK(t *testing.T) {
	r := createServer()

	// borrara el elemento creado en products.json
	req, res := createRequestTest(http.MethodDelete, "/products/1", "")

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
