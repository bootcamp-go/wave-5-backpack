package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Trans_Id          int     `json:"id"`
	Trans_code        string  `json:"code"`
	Trans_currency    string  `json:"currency"`
	Trans_amount      float64 `json:"amount"`
	Trans_transmitter string  `json:"transmitter"`
	Trans_receiver    string  `json:"receiver"`
	Trans_date        string  `json:"date"`
}

type Request struct {
	Trans_Id          int     `json:"id"`
	Trans_code        string  `json:"code"`
	Trans_currency    string  `json:"currency"`
	Trans_amount      float64 `json:"amount"`
	Trans_transmitter string  `json:"transmitter"`
	Trans_receiver    string  `json:"receiver"`
	Trans_date        string  `json:"date"`
}

const FILEPATH string = "./transactions.json"

var T *[]Transaction

func GetAll(ctx *gin.Context) {
	var t Transaction
	transactions, err := readFile(FILEPATH)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, err)
	}
	// ShouldBindQuery setea las variables obtenidas de c.Query("nombredelavariable")
	if err := ctx.ShouldBindQuery(&t); err == nil {
		log.Println(t.Trans_Id)
		log.Println(t.Trans_code)
		log.Println(t.Trans_currency)
		log.Println(t.Trans_amount)
		log.Println(t.Trans_receiver)
		log.Println(t.Trans_transmitter)
		ctx.JSON(http.StatusBadRequest, nil)
	}

	var filtrado []Transaction

	for _, transaction := range transactions {
		if transaction.Trans_Id == t.Trans_Id && transaction.Trans_code == t.Trans_code && transaction.Trans_currency == t.Trans_currency {
			filtrado = append(filtrado, transaction)
		}
	}
	if len(filtrado) > 0 {
		ctx.JSON(http.StatusOK, filtrado)
	} else {
		ctx.JSON(http.StatusOK, transactions)
	}

}

func GetOne(ctx *gin.Context) {
	transactions, err := readFile(FILEPATH)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, err)
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
	found := false
	var t Transaction
	for _, transaction := range transactions {
		if transaction.Trans_Id == id {
			t = transaction
			found = true
			break
		}
	}
	if !found {
		ctx.JSON(http.StatusNotFound, nil)
	} else {
		ctx.JSON(http.StatusOK, t)
	}
}

func Store(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

}

func readFile(path string) ([]Transaction, error) {
	var transactions []Transaction
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if error := json.Unmarshal(data, &transactions); error != nil {
		log.Fatal(error)
		return nil, error
	}
	return transactions, nil
}

func writeFile(file []byte, t Transaction) bool {
	var transactions []Transaction
	if err := json.Unmarshal(file, &transactions); err != nil {
		log.Fatal(err)
		return false
	}
	t.Trans_Id = LastId(transactions)
	transactions = append(transactions, t)
	transactions_json, err := json.Marshal(transactions)
	if err != nil {
		log.Fatal(err)
		return false
	}
	errFile := os.WriteFile(FILEPATH, transactions_json, 0644)
	if errFile != nil {
		log.Fatal(errFile)
	}
	return true
}

func transaction_exist(t []Transaction, id int) bool {
	for _, transaction := range t {
		if transaction.Trans_Id == id {
			return true
		}
	}
	return false
}

func LastId(t []Transaction) int {
	var lastId int = 0
	for i, transaction := range t {
		if t[i].Trans_Id > transaction.Trans_Id {
			lastId = t[i].Trans_Id
		}
	}
	return lastId
}

func main() {
	//creamos un router con Gin
	router := gin.Default()
	tr := router.Group("/transactions")
	tr.POST("/")

	//transacciones
	router.GET("/transactions", GetAll)
	router.GET("/transactions/:id", GetOne)

	//levantamos nuestro servidor en el 8080 por defecto
	router.Run()
}
