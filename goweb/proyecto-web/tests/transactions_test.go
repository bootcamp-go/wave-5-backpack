package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"proyecto-web/cmd/handlers"
	"proyecto-web/internal/domain"
	"proyecto-web/internal/transaction"
	"proyecto-web/pkg/store"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	bd := store.NewStore("../transacciones.json")
	r := transaction.NewRepository(bd)
	service := transaction.NewService(r)
	handler := handlers.NewTransactionHandler(service)
	servidor := gin.Default()

	gr := servidor.Group("transacciones")
	{
		gr.PUT("/:id", handler.Update())
		gr.POST("/", handler.Create())
		gr.GET("/:id", handler.GetById())
		gr.DELETE("/:id", handler.Delete())
	}

	return servidor
}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdate(t *testing.T) {
	//arrange
	transactionToUpdate := domain.Transaction{
		Id:                1,
		CodigoTransaccion: "ACTUALIZADA TEST FUNCIONAL",
		Moneda:            "PESOS ACTUALIZADO",
		Monto:             5.4,
		Emisor:            "SAMSUNG",
		Receptor:          "AFIP",
		FechaTransaccion:  "22-07-2022",
	}

	server := createServer()
	body, _ := json.Marshal(transactionToUpdate)
	request, response := createRequestTest(http.MethodPut, "/transacciones/1", string(body))

	// act
	server.ServeHTTP(response, request)
	var responseStruct struct {
		Code int                `json:"code"`
		Data domain.Transaction `json:"data"`
	}
	err := json.Unmarshal(response.Body.Bytes(), &responseStruct)
	transactionResponse := responseStruct.Data

	//assert
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.Equal(t, "ACTUALIZADA TEST FUNCIONAL", transactionResponse.CodigoTransaccion)
}

func TestDelete(t *testing.T) {
	transactionToDelete := BeforeEachTestDelete(16)
	defer AfterEachTestDelete(transactionToDelete)
	//arrange
	server := createServer()
	request, response := createRequestTest(http.MethodDelete, "/transacciones/16", "")

	// act
	server.ServeHTTP(response, request)
	var responseStruct struct {
		Code int    `json:"code"`
		Data string `json:"data"`
	}
	err := json.Unmarshal(response.Body.Bytes(), &responseStruct)
	transactionResponse := responseStruct.Data

	//assert
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Nil(t, err)
	assert.Equal(t, "Delete exitoso", transactionResponse)
}

func AfterEachTestDelete(transacaction domain.Transaction) {
	server := createServer()
	body, _ := json.Marshal(transacaction)
	request, response := createRequestTest(http.MethodPost, "/transacciones/", string(body))
	server.ServeHTTP(response, request)
}

func BeforeEachTestDelete(id int) domain.Transaction {
	server := createServer()
	request, response := createRequestTest(http.MethodGet, fmt.Sprintf("/transacciones/%d", id), "")
	server.ServeHTTP(response, request)
	var responseStruct struct {
		Code int                `json:"code"`
		Data domain.Transaction `json:"data"`
	}
	json.Unmarshal(response.Body.Bytes(), &responseStruct)
	return responseStruct.Data
}
