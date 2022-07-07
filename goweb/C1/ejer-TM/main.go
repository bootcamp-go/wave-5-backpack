package main

// Crea dentro de la carpeta go-web un archivo llamado main.go
// Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
// Pegale al endpoint para corroborar que la respuesta sea la correcta.

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transacciones struct {
	Id          int     `form:"id"`
	TranCode    string  `form:"tranCode"`
	Currency    string  `form:"currency"`
	Ammount     float64 `form:"ammount"`
	Transmitter string  `form:"transmitter"`
	Reciever    string  `form:"reciever"`
	TranDate    string  `form:"tranDate"`
}

func Read() (*[]Transacciones, error) {
	data, err := os.ReadFile("./transacciones.json")
	if err != nil {
		return nil, errors.New("error al leer el archivo")
	}

	var t []Transacciones
	if err := json.Unmarshal(data, &t); err != nil {
		fmt.Println("error aqui")
		log.Fatal(err)
	}

	return &t, nil
}

func GetAllTransactions(ctx *gin.Context) {
	transactions, err := Read()
	if err != nil {
		ctx.JSON(500, gin.H{"Error:": err.Error()})
		return
	}
	for _, transaction := range *transactions {
		ctx.JSON(200, transaction)
	}

}

func SearchByIdHandler(ctx *gin.Context) {
	transactions, err := Read()
	id := ctx.Param("id")

	if err != nil {
		ctx.JSON(500, gin.H{"Error:": err.Error()})
		return
	}

	for _, transaction := range *transactions {

		idT := strconv.Itoa(transaction.Id)
		if idT == id {
			ctx.JSON(200, transaction)
			return
		}
	}
	ctx.JSON(404, gin.H{"Error": "Id not found"})
}

func FiltrarTransactionsHandler(ctx *gin.Context) {
	transactions, err := Read()

	if err != nil {
		ctx.JSON(500, gin.H{"Error:": err.Error()})
		return
	}
	var tr Transacciones

	if err := ctx.ShouldBindQuery(&tr); err != nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
		ctx.JSON(400, err)
		return
	}

	var filter []*Transacciones
	for _, t := range *transactions { // filtrado por todos los campos
		if tr.Id == t.Id {
			filter = append(filter, &t)
		}
	}
	// && tr.TranCode == t.TranCode && tr.Currency == t.Currency && tr.Transmitter == t.Transmitter && tr.Reciever == t.Reciever && tr.TranDate == t.TranDate

	ctx.JSON(http.StatusOK, filter) // devovemos el array filtrado
}

func main() {

	// const name string = "Seba"

	// data, _ := os.ReadFile("./transacciones.json")

	// var t []transacciones

	// if err := json.Unmarshal(data, &t); err != nil {
	// 	fmt.Println("error aqui")
	// 	log.Fatal(err)
	// }

	// router.GET("/saludo", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hola " + name,
	// 	})
	// })

	router := gin.Default()

	router.GET("/transacciones", GetAllTransactions)

	router.GET("transacciones/:id", SearchByIdHandler)

	router.GET("/filtrartransaction", FiltrarTransactionsHandler)

	router.Run()

}
