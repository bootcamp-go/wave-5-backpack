package test

import (
	"bytes"
	"clase4-testing-prt1/cmd/server/handler"
	"clase4-testing-prt1/internal/transactions"
	"clase4-testing-prt1/pkg/bank"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := bank.New(bank.FileType, pathDB)
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	p := handler.NewTransaction(service)

	r := gin.Default()

	pr := r.Group("/transactions")
	// pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	// pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())
	pr.PUT("/:id", p.Update())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllTransactions(t *testing.T) {

	type transaction struct {
		ID                int     `json:"id" binding:"-"`
		CodigoTransaccion string  `json:"codigo de transaccion"`
		Moneda            string  `json:"moneda"`
		Monto             float64 `json:"monto"`
		Emisor            string  `json:"emisor"`
		Receptor          string  `json:"receptor"`
		Fecha             string  `json:"fecha de transaccion"`
	}

	// crear el Server y definir las Rutas
	r := createServer("transactions.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	var objRes []transaction

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}

func TestUpdateTransaction(t *testing.T) {
	r := createServer("transaction_update.json")

	req, rr := createRequestTest(http.MethodPut, "/transactions/3",
		`{
			"codigo de transaccion": "jh1kjna",
			"moneda": "USD ",
			"monto": 8173.76,
			"emisor": "Pancho",
			"receptor": "Gabriel",
			"fecha de transaccion": "2020-09-12"
		  }`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateDelete(t *testing.T) {
	r := createServer("transaction_update.json")

	req, rr := createRequestTest(http.MethodDelete, "/transactions/3", ``)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
