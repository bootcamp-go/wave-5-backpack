package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transaction struct {
	Id        int     `json:"-"`
	Code      string  `json:"code" binding:"required"`
	Currency  string  `json:"currency" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Issuer    string  `json:"issuer" binding:"required"`
	Recipient string  `json:"recipient" binding:"required"`
	Date      string  `json:"date" binding:"required"`
}

var transactions []transaction

func openFile() {
	file, _ := os.ReadFile("transactions.json")
	json.Unmarshal([]byte(file), &transactions)
}

func GetAll(ctx *gin.Context) {
	var ans []transaction

	issuer := ctx.Query("issuer")
	date := ctx.Query("date")

	for _, transaction := range transactions {
		filter := true
		if issuer != "" && issuer != transaction.Issuer {
			filter = false
		}
		if date != "" && date != transaction.Date {
			filter = false
		}
		if filter {
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
	for _, transaction := range transactions {
		if strconv.Itoa(transaction.Id) == ctx.Param("id") {
			ctx.JSON(200, transaction)
			return
		}
	}
	ctx.JSON(404, "Id no encontrado")
}

func getLastId() int {
	if len(transactions) == 0 {
		return 0
	} else {
		return transactions[len(transactions)-1].Id
	}
}

func validateFields(err *error) string {
	errAns := ""
	var ve validator.ValidationErrors
	if errors.As(*err, &ve) {
		for _, fe := range ve {
			errAns += fmt.Sprintf("el campo %s es requerido ", fe.Field())
		}
	}
	return errAns
}

func Create(ctx *gin.Context) {
	if ctx.GetHeader("token") != "elpepe" {
		ctx.JSON(401, gin.H{
			"error": "token inválido",
		})
		return
	}

	var req transaction
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"Error: ": validateFields(&err)})
		return
	}
	newId := getLastId() + 1
	req.Id = newId
	transactions = append(transactions, req)
	ctx.JSON(200, req)
}

func main() {
	openFile()
	router := gin.Default()
	router.GET("/transactions", GetAll)
	router.GET("/transaction/:id", GetOne)
	router.POST("/transaction", Create)
	router.Run()
}
