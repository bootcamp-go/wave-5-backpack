package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewStore("../users.json")
	repo := users.NewRepository(db)
	servi := users.NewService(repo)
	u := handler.NewUser(servi)
	router := gin.Default()
	usr := router.Group("/users")
	usr.PUT("/:id", u.Update())
	usr.DELETE("/:id", u.Delete())
	usr.GET("/", u.GetAll())
	return router
}
func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", "123456")
	return req, httptest.NewRecorder()
}

func TestUpdateUser(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/users/1", `{
        "name":"new","last_name":"new","email":"b@b.com","age":10,"height":1.20,"active":true,"creation_date":"new"
    }`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestDeleteUser(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/users/5", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 204, rr.Code)
}

func TestGetAllUsers(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodGet, "/users/", "")
	var resp web.Response
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.True(t, reflect.ValueOf(resp.Data).Len() > 0)
}
