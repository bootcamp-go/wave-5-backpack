package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	Id        int
	Code      string
	Currency  string
	Amount    float64
	Issuer    string
	Recipient string
	Date      string
}

func GetAll(ctx *gin.Context) {
	file, _ := os.ReadFile("transactions.json")
	var transactions, ans []transaction
	json.Unmarshal([]byte(file), &transactions)
	issuer := ctx.Query("issuer")
	date := ctx.Query("date")

	filter := make([]bool, len(transactions))
	for i, _ := range filter {
		filter[i] = true
	}
	for i, transaction := range transactions {
		if issuer != "" && issuer != transaction.Issuer {
			filter[i] = false
		}
		if date != "" && date != transaction.Date {
			filter[i] = false
		}
		if filter[i] {
			ans = append(ans, transaction)
		}
	}
	if len(ans) == 0 {
		ctx.JSON(404, "Nada fue encontrado")
		return
	}
	ctx.JSON(200, ans)
}

func GetOne(ctx *gin.Context) {
	file, _ := os.ReadFile("transactions.json")
	var transactions []transaction
	json.Unmarshal([]byte(file), &transactions)

	for _, transaction := range transactions {
		if strconv.Itoa(transaction.Id) == ctx.Param("id") {
			ctx.JSON(200, transaction)
			return
		}
	}
	ctx.JSON(404, "Id no encontrado")
}

func main() {
	router := gin.Default()
	router.GET("/transactions", GetAll)
	router.GET("/transaction/:id", GetOne)
	router.Run()
}
