package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var transactions []Transaction

const SECRET_TOKEN = "1234567"

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

type Transaction struct {
	Id              int       `json:"id" binding:"-"`
	TransactionCode string    `json:"transaction_code" binding:"-"`
	Currency        string    `json:"currency" binding:"required"`
	Amount          float64   `json:"amount" binding:"required"`
	Sender          string    `json:"sender" binding:"required"`
	Reciever        string    `json:"reciever" binding:"required"`
	TransactionDate time.Time `json:"transaction_date" binding:"-"`
}

func JSONTransactionToStruct(jsonBytes []byte) (*[]Transaction, error) {
	var transactions []Transaction
	err := json.Unmarshal(jsonBytes, &transactions)
	return &transactions, err
}

type FilterFunction (func() bool)

func containsString(base, target string) FilterFunction {
	return func() bool {
		return target == "" || strings.Contains(base, target)
	}

}

func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func sameDay(date time.Time, dateTarget string) FilterFunction {
	return func() bool {
		if dateTarget == "" {
			return true
		}
		dateParsed, err := time.Parse("02-01-2006", dateTarget)
		if err != nil {
			return false
		}
		return DateEqual(date, dateParsed)
	}
}

func eqString(base, target string) FilterFunction {
	return func() bool {
		return target == "" || target == base
	}

}

func eqAmount(base float64, target string) FilterFunction {
	return func() bool {
		value, err := strconv.ParseFloat(target, 64)
		if target == "" {
			return true
		}
		if err != nil {
			return false
		}
		return value == base
	}
}

func maxAmount(base float64, target string) FilterFunction {
	return func() bool {
		value, err := strconv.ParseFloat(target, 64)
		if target == "" {
			return true
		}
		if err != nil {
			return false
		}
		return value >= base
	}
}

func minAmount(base float64, target string) FilterFunction {
	return func() bool {
		value, err := strconv.ParseFloat(target, 64)
		if target == "" {
			return true
		}
		if err != nil {
			return false
		}
		return value <= base
	}
}

func passFilters(filters ...FilterFunction) bool {

	flagFilter := true
	for _, filter := range filters {
		if !filter() {
			flagFilter = false
			break
		}
	}
	return flagFilter
}

func Search(ctx *gin.Context) {

	var filteredTransactions = make([]Transaction, 0)
	for _, transaction := range transactions {
		if result := passFilters(
			containsString(transaction.Sender, ctx.Query("sender")),
			containsString(transaction.Reciever, ctx.Query("reciever")),
			eqString(transaction.Currency, ctx.Query("currency")),
			eqAmount(transaction.Amount, ctx.Query("amount")),
			minAmount(transaction.Amount, ctx.Query("minAmount")),
			maxAmount(transaction.Amount, ctx.Query("maxAmount")),
			sameDay(transaction.TransactionDate, ctx.Query("date")),
		); result {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}

	ctx.JSON(200, gin.H{
		"data": filteredTransactions,
	})
}

func GetAll(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data": transactions,
	})
}

func GetOne(ctx *gin.Context) {

	searchId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "param id must be integer",
		})
		return
	}
	var transactionResult *Transaction
	for _, transaction := range transactions {
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

func CreateTransaction() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != SECRET_TOKEN {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Access Denied: Token Unauthorized",
			})
			return
		}
		transaction := Transaction{}
		if err := ctx.ShouldBindJSON(&transaction); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				errors := err.(validator.ValidationErrors)
				invalidErrors := make([]string, 0)
				for _, e := range errors {
					field := e.Field()
					invalidErrors = append(invalidErrors, fmt.Sprintf("el campo %s es requerido",
						field))
				}
				ctx.JSON(400, gin.H{
					"errors": invalidErrors,
				})
			} else {
				ctx.JSON(400, gin.H{
					"error": err.Error(),
				})
			}

			return
		}
		if transaction.Amount <= 0 {
			ctx.JSON(400, gin.H{
				"error": "el campo amount debe ser mayor que 0",
			})
			return
		}
		nextId := transactions[len(transactions)-1].Id + 1
		transaction.Id = nextId
		transaction.TransactionDate = time.Now()
		transaction.TransactionCode = randomString(30)
		transactions = append(transactions, transaction)
		ctx.JSON(201, transaction)
	}

}

func main() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	router := gin.Default()

	transactionsBytes, err := os.ReadFile("./transactions.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	transactionsJSON, err := JSONTransactionToStruct(transactionsBytes)
	transactions = *transactionsJSON
	if err != nil {
		fmt.Println(err)
		return
	}

	transactionsRoute := router.Group("/transactions")
	transactionsRoute.GET("/", GetAll)
	transactionsRoute.GET("/search", Search)
	transactionsRoute.GET("/:id", GetOne)
	transactionsRoute.POST("/", CreateTransaction())
	router.Run()
}
