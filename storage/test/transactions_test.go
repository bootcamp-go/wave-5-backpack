package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestTransaction struct {
	Id              int       `json:"id" binding:"-"`
	TransactionCode string    `json:"transaction_code" binding:"-"`
	Currency        string    `json:"currency" binding:"required"`
	Amount          float64   `json:"amount" binding:"required"`
	Sender          string    `json:"sender" binding:"required"`
	Reciever        string    `json:"reciever" binding:"required"`
	TransactionDate time.Time `json:"transaction_date" binding:"-"`
}

type ResponseUpdateTransaction struct {
	Code  int             `json:"code"`
	Data  TestTransaction `json:"data,omitempty"`
	Error string          `json:"err,omitempty"`
}

type ResponseDeleteTransaction struct {
	Code  int    `json:"code"`
	Data  string `json:"data,omitempty"`
	Error string `json:"err,omitempty"`
}

func TestUpdateTransactionOK(t *testing.T) {
	r := createServer()
	payload := fmt.Sprintf("{\"currency\": \"%s\", \"amount\": %f, \"sender\": \"%s\", \"reciever\": \"%s\"}",
		"ARG",
		15000.00,
		"Anonimo",
		"Anonimo2",
	)
	req, rr := createRequestTest(http.MethodPut, "/transactions/2", payload)

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
	var resData ResponseUpdateTransaction
	err := json.Unmarshal(rr.Body.Bytes(), &resData)

	assert.Nil(t, err)
	assert.Equal(t, resData.Data.Id, 2)
	assert.Equal(t, resData.Data.Amount, 15000.00)
	assert.Equal(t, resData.Data.Currency, "ARG")
	assert.Equal(t, resData.Data.Sender, "Anonimo")
	assert.Equal(t, resData.Data.Reciever, "Anonimo2")
}

func TestDeleteTransactionOK(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "/transactions/2", "")

	r.ServeHTTP(rr, req)
	assert.Equal(t, 202, rr.Code)
	var resData ResponseDeleteTransaction
	err := json.Unmarshal(rr.Body.Bytes(), &resData)

	assert.Nil(t, err)
	expectedResponse := fmt.Sprintf("transaction with id %d was deleted successfully", 2)
	assert.Equal(t, expectedResponse, resData.Data)
}
