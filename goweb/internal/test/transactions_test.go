package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/pkg/storage"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type request struct {
	Monto    float64 `json:"monto"`
	Cod      string  `json:"cod_transaction"`
	Moneda   string  `json:"moneda"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
}

// levanta un servidor para testear, las capas y agrega las rutas
func createServer(t *testing.T) *gin.Engine {
	err := godotenv.Load()
	assert.Nil(t, err)

	// Init capas testing
	storage := storage.NewStorage("transactions.json")
	repo := transactions.NewRepository(storage)
	service := transactions.NewService(repo)
	tr := handler.NewTransaction(service)

	//Engine
	r := gin.Default()

	// Middleware
	r.Use(web.TokenAuthMiddleware())

	// Router
	rt := r.Group("/transactions")
	{
		rt.GET("", tr.GetAll)
		rt.GET("/:id", tr.GetByID)

		rt.PUT("/:id", tr.Update)
		rt.PATCH("/:id", tr.Patch)

		rt.POST("", tr.CreateTransaction)

		rt.DELETE("/:id", tr.Delete)
	}

	return r
}

// Retorna la request y un ResponseRecorder donde se guardará la response
func createRequest(method, url string, data []byte) (*http.Request, *httptest.ResponseRecorder) {
	bodyReq := bytes.NewReader(data) // Reader para NewRequest

	req := httptest.NewRequest(method, url, bodyReq)
	req.Header.Add("token", "1245")

	return req, httptest.NewRecorder()
}

func TestFunctionalUpdate(t *testing.T) {
	//Arrange
	bodyReq := request{
		Monto:    1000,
		Cod:      "update cod",
		Moneda:   "USD",
		Emisor:   "BBVA",
		Receptor: "Mercado Pago",
	}

	data, err := json.Marshal(bodyReq)
	assert.Nil(t, err)

	server := createServer(t)
	req, rr := createRequest(http.MethodPut, "/transactions/1", data)

	//Act
	server.ServeHTTP(rr, req)

	//Assert
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestFunctionalDelete(t *testing.T) {
	//Arrange
	server := createServer(t)
	req, rr := createRequest(http.MethodDelete, "/transactions/2", nil)

	//Act
	server.ServeHTTP(rr, req)

	//Assert
	assert.Equal(t, http.StatusOK, rr.Code)
}
