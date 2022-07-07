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
	"github.com/go-playground/validator/v10"
)

type Transacciones struct {
	Id          int     `form:"id"`
	TranCode    string  `json:"tranCode" binding:"required"`
	Currency    string  `json:"currency" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Transmitter string  `json:"transmitter" binding:"required"`
	Reciever    string  `json:"reciever" binding:"required"`
	TranDate    string  `json:"tranDate" binding:"required"`
}

var t []Transacciones
var lastID int

func Read() (*[]Transacciones, error) {
	data, err := os.ReadFile("./transacciones.json")
	if err != nil {
		return nil, errors.New("error al leer el archivo")
	}

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

func GenerateTransaction() gin.HandlerFunc {
	var transac Transacciones
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.JSON(401, gin.H{
				"error": "No tiene permisos para realizar la peticion solicitada",
			})
			return
		}
		if err := c.ShouldBindJSON(&transac); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("El campo %s es requerido y ", field.Field())
					} else {
						result += fmt.Sprintf("El campo %s es requerido", field.Field())
					}
				}
				c.JSON(404, result)
				return
			}
		}
		lastID++
		transac.Id = lastID
		t = append(t, transac)
		c.JSON(200, t)
	}
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

	router.POST("transactions", GenerateTransaction())

	// tr := router.Group("/transacciones")
	// tr.POST("/", GenerateTransaction())

	router.Run()

}
