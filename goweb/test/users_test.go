package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "62c5b68a0cc23a33375c85f8")

	db := store.NewStore(pathDB)
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	user := handler.NewUser(service)

	router := gin.Default()

	u := router.Group("/users")
	{
		u.PUT("/:id", user.Update())
	}
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "62c5b68a0cc23a33375c85f8")

	return req, httptest.NewRecorder()
}

func Test_UpdateUser_OK(t *testing.T) {
	r := createServer("users_test.json")

	req, rr := createRequestTest(http.MethodPut, "/users/1",
		`{"nombre": "Daniela", "apellido": "Bedoya", "email": "djfsj@gmail.com", "edad": 20, "altura": 1.45, "activo": true, "fechaCreacion": "2021-10-02T04:44:12 +05:00" }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
