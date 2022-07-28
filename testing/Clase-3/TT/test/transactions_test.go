package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "12345")
	db := store.NewStore("./transactions_test.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	r := gin.Default()

	pr := r.Group("/transactions")
	pr.POST("/", t.Store())
	pr.GET("/", t.GetAll())
	pr.PUT("/:id", t.Update())
	pr.DELETE("/:id", t.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}
func TestGetAll(t *testing.T) {

	r := createServer()

	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	var objRes web.Response

	r.ServeHTTP(rr, req)

	err := json.Unmarshal(rr.Body.Bytes(), &objRes)

	assert.Nil(t, err)
	assert.True(t, reflect.ValueOf(objRes.Data).Len() > 0)
	assert.Equal(t, 200, rr.Code)
}

func TestUpdate(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/transactions/3",
		`{"tranCode":"XXX1234","currency":"CLP","amount":300,"transmitter":"MERCADOPAGO","receiver":"PEDRO","tranDate":"28-07-22"}`)

	r.ServeHTTP(rr, req)
	fmt.Println(rr.Code)
	assert.Equal(t, http.StatusAccepted, rr.Code)
}

func TestDelete(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/transactions/1", "")

	r.ServeHTTP(rr, req)
	fmt.Println(rr.Code)

	assert.Equal(t, http.StatusAccepted, rr.Code)
}

func TestPost(t *testing.T) {

	r := createServer()
	req, rr := createRequestTest(http.MethodPost, "/transactions/",
		`{"tranCode":"XXX1234","currency":"CLP","amount":300,"transmitter":"MERCADOPAGO","receiver":"JUAN","tranDate":"28-julio-22"}`)

	r.ServeHTTP(rr, req)
	fmt.Println(rr.Code)

	assert.Equal(t, 200, rr.Code)
}
