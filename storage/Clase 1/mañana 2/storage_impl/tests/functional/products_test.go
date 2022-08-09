package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bootcamp-go/storage/cmd/server/handler"
	cnn "github.com/bootcamp-go/storage/db"
	"github.com/bootcamp-go/storage/internal/domains"
	"github.com/bootcamp-go/storage/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var s = createServer()

func createServer() *gin.Engine {
	os.Setenv("USERNAME", "root")
	os.Setenv("PASSWORD", "root")
	os.Setenv("DATABASE", "storage")

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)

	p := handler.NewProduct(serv)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetByName())
	pr.POST("/", p.Store())

	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestStoreProduct_Ok(t *testing.T) {
	new := domains.Product{
		Name:  "producto nuevo",
		Type:  "producto tipo",
		Count: 3,
		Price: 84.4,
	}

	product, err := json.Marshal(new)
	require.Nil(t, err)

	req, rr := createRequest(http.MethodPost, "/api/v1/products/", string(product))
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// struct for assertion
	p := struct{ Data domains.Product }{}
	err = json.Unmarshal(rr.Body.Bytes(), &p)
	require.Nil(t, err)

	new.ID = p.Data.ID
	assert.Equal(t, new, p.Data)
}

func TestGetByNameProduct_Ok(t *testing.T) {
	req, rr := createRequest(http.MethodGet, "/api/v1/products/", `{"nombre":"producto nuevo"}`)
	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAll_Ok(t *testing.T){
	req, rr := createRequest(http.MethodGet, "api/v1/products/","")
	s.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

