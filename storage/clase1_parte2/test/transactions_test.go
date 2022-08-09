package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"goweb/pkg/web"
)

func createServer(pathDB string) *gin.Engine {
	//Simulo pasarle un TOKEN
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewStore(pathDB)
	repositoryTransaction := transactions.NewRepository(db)
	serviceTransaction := transactions.NewService(repositoryTransaction)
	handlerTransaction := handler.NewHandler(serviceTransaction)

	r := gin.Default()

	tg := r.Group("/transactions")
	tg.GET("/", handlerTransaction.GetAll())
	tg.POST("/", handlerTransaction.Store())
	tg.PUT("/:id", handlerTransaction.Upddate())
	tg.DELETE("/:id", handlerTransaction.Delete())

	return r
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	//Realiza la misma funcion que un POSTMAN, crea una ambtio request para probar los endpoints
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllTransactions(t *testing.T) {

	//Creo el server y defino rutas
	r := createServer("./transactions.json")

	//Creo el request de tipo POST y response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/transactions/", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	//Defino el objeto Response
	var objectResponse web.Response
	//var objectResponse interface{}

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objectResponse)
	assert.Nil(t, err)
	assert.True(t, reflect.ValueOf(objectResponse.Data).Len() > 0)
}

func TestSaveTransaction(t *testing.T) {

	//Creo el server y defino rutas
	r := createServer("./transactions.json")

	//Creo el request de tipo POST y response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/transactions/", `{
        "CodTransaction": "123zwy","Currency": "COP","Amount": 99999,"Sender": "Valentina", "Receiver": "Michael", "DateOrder": "10-05-2022"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestUpdateTransaction(t *testing.T) {

	//Creo el server y defino rutas
	r := createServer("./transactions.json")

	//Creo el request de tipo POST y response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/transactions/1002", `{
        "cod_transaction": "32190812","Currency": "EUR","Amount": 123456,"Sender": "Vale", "Receiver": "Michael", "date_order": "10-12-23"
    }`)

	//Indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	//Defino el objeto Response
	var objectResponse web.Response
	//var objectResponse interface{}

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objectResponse)
	assert.Nil(t, err)
}

func TestDeleteTransaction(t *testing.T) {

	//Creo el server y defino rutas
	r := createServer("./transactions.json")

	//Creo el request de tipo POST y response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/transactions/1001", `{}`)

	//Indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}