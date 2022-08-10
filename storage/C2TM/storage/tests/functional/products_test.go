package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"storage/cmd/server/handler"

	cnn "storage/db"

	"storage/internal/domain"
	"storage/internal/products"

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
	// pr.GET("/", p.GetByName())
	pr.POST("/", p.Store())

	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestStoreProduct_Ok(t *testing.T) {
	new := domain.Product{
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
	p := struct{ Data domain.Product }{}
	err = json.Unmarshal(rr.Body.Bytes(), &p)
	require.Nil(t, err)

	new.ID = p.Data.ID
	assert.Equal(t, new, p.Data)
}

func TestGetByName(t *testing.T) {
	var dataSource = "root@tcp(localhost:3306)/storageC1TT"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	myRepo := products.NewRepository(StorageDB)

	product, err := myRepo.GetByName("Computador")
	assert.Nil(t, err)
	assert.Equal(t, "Computador", product.Name)
}

func TestGetAll(t *testing.T) {
	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	products, err := repo.GetAll()

	assert.NoError(t, err)
	assert.True(t, len(products) > 0)
}
