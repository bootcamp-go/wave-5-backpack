package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transactions struct {
	Id                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_de_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_de_transaccion"`
}

var T []Transactions

func main() {

	data, err := os.ReadFile("transactions.json")
	if err != nil {
		fmt.Println(err)
	}

	if err := json.Unmarshal(data, &T); err != nil {
		log.Fatal(err)
	}
	router := gin.Default()

	router.GET("/Hola_gopher", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello amiguitos!",
		})
	})

	router.GET("/transactions", GetAll)
	router.GET("/transactions/:id", FindTransactionbyID)
	router.GET("/transactions/field", FindTransactionbyField) // waiting a GET petition with query
	router.Run()
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": T,
	})
}

// GetOne
func FindTransactionbyID(ctx *gin.Context) {
	var TransactionbyId Transactions
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(404, "parameter fail is not an Id")
	}
	if id < len(T) {
		TransactionbyId = T[id]
		ctx.JSON(200, gin.H{
			"message": TransactionbyId,
		})
	} else {
		ctx.String(404, "transaction not found")
	}

}
func FindTransactionbyField(ctx *gin.Context) {
	fmt.Println("find transactions")
	ctx.String(200, "donde esta el Juanchi?")
	query := ctx.Request.URL.Query()

	fmt.Println(query)

}

// func FindId(id string, t []Transactions) Transactions{

// 	return t[1]
// }
