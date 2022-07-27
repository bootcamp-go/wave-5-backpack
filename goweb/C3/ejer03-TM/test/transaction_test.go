package test

import (
	"bytes"
	"ejer02-TT/cmd/server/handler"
	"ejer02-TT/internal/transactions"
	"ejer02-TT/pkg/store"
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

	db := store.NewStore("transacciones.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	p := handler.NewTransaction(service)

	r := gin.Default()

	pr := r.Group("/transactions")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PATCH("/:id", p.UpdateCodeAndAmount())
	//pr.DELETE("/:id", p.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}

func TestGetAllProducts(t *testing.T) {

	type transaccion struct {
		Id          int     `json:"id"`
		TranCode    string  `json:"tranCode"`
		Currency    string  `json:"currency"`
		Amount      float64 `json:"amount"`
		Transmitter string  `json:"transmitter"`
		Reciever    string  `json:"reciever"`
		TranDate    string  `json:"tranDate"`
	}

	// crear el Server y definir las Rutas
	r := createServer("transacciones.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	var objRes []transaccion

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}

func TestSaveProduct(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("transacciones.json")
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
        "tranCode": "code",
		"currency": "curr",
		"amount": 11,
		"transmitter":"prueba", 
		"reciever": "rec", 
		"tranDate": "date"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

// func TestUpdateCodeAndAmmount(t *testing.T) {
// 	r := createServer("transacciones.json")

// 	req, rr := createRequestTest(http.MethodPatch, "/transactions/3",
// 		`{"tranCode": "1", "amount": 2}`)

// 	r.ServeHTTP(rr, req)

// 	assert.Equal(t, 200, rr.Code)
// }
