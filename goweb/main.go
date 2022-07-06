package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	if currencyType := ctx.Query("currency"); currencyType != "" {
		currencyType = strings.ToUpper(currencyType)
		var filteredTransactions = make([]Transaction, 0)
		for _, transaction := range *transactions {
			if transaction.Currency == currencyType {
				filteredTransactions = append(filteredTransactions, transaction)
			}

		}
		transactions = &filteredTransactions
	}
	ctx.JSON(200, gin.H{
		"data": transactions,
	})
}

func GetOne(ctx *gin.Context) {

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
	searchId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "param id must be integer",
		})
		return
	}
	var transactionResult *Transaction
	for _, transaction := range *transactions {
		if transaction.Id == searchId {
			transactionResult = &transaction
			break
		}
	}
	if transactionResult == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err": "transaction not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": transactionResult,
	})
}

func main() {
	router := gin.Default()

	transactionsRoute := router.Group("/transactions")
	transactionsRoute.GET("/", GetAll)
	transactionsRoute.GET("/:id", GetOne)

	router.Run()
}
