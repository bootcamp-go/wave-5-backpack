package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id              int       `json:"id"`
	TransactionCode string    `json:"transaction_code"`
	Currency        string    `json:"currency"`
	Amount          float64   `json:"amount"`
	Sender          string    `json:"sender"`
	Reciever        string    `json:"reciever"`
	TransactionDate time.Time `json:"transaction_date"`
}

func JSONTransactionToStruct(jsonBytes []byte) (*[]Transaction, error) {
	var transactions []Transaction
	err := json.Unmarshal(jsonBytes, &transactions)
	return &transactions, err
}

func GetAll(ctx *gin.Context) {
	transactionsBytes, err := os.ReadFile("./transactions.json")
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}

	transactions, err := JSONTransactionToStruct(transactionsBytes)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"err": "bad json file : " + err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"data": transactions,
	})
}

func main() {
	router := gin.Default()

	router.GET("/transactions", GetAll)

	router.Run()
}
