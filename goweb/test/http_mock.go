package test

import (
	"bytes"
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "1234567")
	db := store.NewStore("./transactions_test.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)
	r := gin.Default()

	tr := r.Group("/transactions")
	tr.PUT("/:id", t.Update())
	tr.DELETE("/:id", t.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", "1234567")

	return req, httptest.NewRecorder()
}
