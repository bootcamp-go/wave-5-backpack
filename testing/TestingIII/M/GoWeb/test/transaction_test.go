package test

import (
	"GoWeb/cmd/server/handler"
	"GoWeb/internals/transactions"
	"GoWeb/pkg/store"
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	//solicitud y validacion de token
	err := os.Setenv("TOKEN", "12345")

	if err != nil {
		panic("error al codificar el Token")
	}

	//instanciando capas
	db := store.New(store.FileType, pathDB)
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	p := handler.NewTransaction(service)

	//servidor por defecto
	r := gin.Default()
	//verbos http a utilizar
	pr := r.Group("/transacciones")
	pr.PATCH("/:id", p.UpdateCode())
	pr.DELETE("/:id", p.Delete())

	return r
}

func createrRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}

func TestUpdateCode(t *testing.T) {
	r := createServer("update.json")

	req, rr := createrRequestTest(http.MethodPatch, "/transacciones/1",
		`{
		"code":"AAA",
		"amount":2000}`)
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDelete(t *testing.T) {
	r := createServer("delete.json")

	req, rr := createrRequestTest(http.MethodDelete, "/transacciones/1",
		`{}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}
