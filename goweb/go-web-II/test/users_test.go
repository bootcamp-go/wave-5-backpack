package test

import (
	"bytes"
	"goweb/go-web-II/cmd/handler"
	"goweb/go-web-II/internal/products"
	"goweb/go-web-II/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "1234")
	db := store.NewStore(pathDB)
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()

	ur := r.Group("/users")
	ur.PUT("/:id", u.Update())
	ur.DELETE("/:id", u.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1234")

	return req, httptest.NewRecorder()
}

func TestUpdateUser(t *testing.T) {
	r := createServer("../usuarios.json")
	req, rr := createRequestTest(http.MethodPut, "/users/1", `{
		"name": "Nahuelito", "surname": "Rodriguez", "email": "prueba@gmail.com", "age": 20, "active": true, "created": "24/11/2009"
	}`)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestDeleteUser(t *testing.T) {
	r := createServer("../usuarios.json")
	req, rr := createRequestTest(http.MethodDelete, "/users/1", "")

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}
