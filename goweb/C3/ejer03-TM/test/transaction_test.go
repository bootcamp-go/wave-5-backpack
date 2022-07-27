package test

import (
	"bytes"
	"ejer02-TT/cmd/server/handler"
	"ejer02-TT/internal/transactions"
	"ejer02-TT/pkg/store"
	"ejer02-TT/pkg/web"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "12345")

	db := store.NewStore("transacciones.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	tr := r.Group("/transactions")
	tr.POST("/", t.Store())
	tr.GET("/", t.GetAll())
	tr.PATCH("/:id", t.UpdateCodeAndAmount())
	//pr.DELETE("/:id", p.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}

func TestGetAllTransactions(t *testing.T) {

	// crear el Server y definir las Rutas
	r := createServer("transacciones.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	var objRes web.Response

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, reflect.ValueOf(objRes.Data).Len() > 0)
}

func TestSaveTransaction(t *testing.T) {
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
