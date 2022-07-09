package main

import (
	"github.com/gin-gonic/gin"
)

type Transactions struct {
	Id              int     `json:"id" binding:"required"`
	TransactionCode string  `json:"transaction_code"binding:"required"`
	TypeCurrency    string  `json:"type_of_currency"binding:"required"`
	Amount          float64 `json:"amount"binding:"required"`
	Transmitter     string  `json:"transmitter"binding:"required"`
	Receiver        string  `json:"receiver"binding:"required"`
	Date            string  `json:"date"binding:"required"`
	Completed       bool    `json:"completed"binding:"required"`
}

var Req Transactions

func PostTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&Req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		Req.Id++
		c.JSON(200, Req)
	}
}
