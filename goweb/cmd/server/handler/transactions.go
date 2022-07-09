package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/internal/transactions"
	"net/http"
	"strconv"
)

type Request struct {
	Id              int64   `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	TypeCurrency    string  `json:"type_of_currency"`
	Amount          float64 `json:"amount"`
	Transmitter     string  `json:"transmitter"`
	Receiver        string  `json:"receiver"`
	Date            string  `json:"date"`
	Completed       bool    `json:"completed"`
}

type Transactions struct {
	service transactions.Service
}

func NewTransaction(s transactions.Service) *Transactions {
	return &Transactions{service: s}
}

func (t *Transactions) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}

		transactions, err := t.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token not valid"})
			return
		}

		if len(transactions) <= 0 {
			c.JSON(http.StatusOK, gin.H{"message": "No transactions available."}) // 500
			return
		}

		c.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}

func (t *Transactions) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"}) //401
			return
		}

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //400
			return
		}

		transaction, err := t.service.Store(req.TransactionCode, req.TypeCurrency, req.Transmitter, req.Receiver, req.Date, req.Amount, req.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //500
			return
		}

		c.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

func (t *Transactions) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token not valid",
			})
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if req.TransactionCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Transaction Code is Required",
			})
			return
		}
		if req.TypeCurrency == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Type of Currency is Required",
			})
			return
		}
		if req.Amount <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Amount is Required",
			})
			return
		}
		if req.Transmitter == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Transmitter Field is Required",
			})
			return
		}
		if req.Receiver == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Receiver Field is Required",
			})
			return
		}
		if req.Date == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Date is Required",
			})
			return
		}
		u, err := t.service.Update(int64(id), req.TransactionCode, req.TypeCurrency, req.Date, req.Transmitter, req.Receiver, req.Amount, req.Completed)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, u)
	}
}

func (t Transactions) UpdateTransmitter() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id format invalid"})
			return
		}
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Transmitter == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Transmitter field is Required"})
			return
		}
		p, err := t.service.UpdateTransmitter(id, req.Transmitter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (t Transactions) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not valid"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id format invalid"})
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Product with id-%d has been deleted ", id))
	}
}
