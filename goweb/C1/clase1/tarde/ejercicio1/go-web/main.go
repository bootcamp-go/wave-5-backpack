package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 1 - Filtremos nuestro endpoint

Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo
se tiene que poder filtrar por todos los campos.
	1. Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
	2. Luego genera la lógica de filtrado de nuestro array.
	3. Devolver por el endpoint el array filtrado.
*/

type Transaction struct {
	Id                int64   `form:"id"`
	CodigoTransaccion string  `form:"codigo"`
	Moneda            string  `form:"moneda"`
	Monto             float64 `form:"monto"`
	Emisor            string  `form:"emisor"`
	Receptor          string  `form:"receptor"`
}

func main() {
	// Transaciones
	router := gin.Default()
	router.GET("/filtrartransaction", FiltrarTransactionsHandler)

	router.Run(":8000") // Por defecto gin arranca el server on 8080
}

func FiltrarTransactionsHandler(c *gin.Context) {
	listaTransactions := generarTransactions()
	var tt Transaction
	if c.ShouldBindQuery(&tt) == nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
		log.Println(tt.Id)
		log.Println(tt.CodigoTransaccion)
		log.Println(tt.Moneda)
		log.Println(tt.Monto)
		log.Println(tt.Emisor)
		log.Println(tt.Receptor)
	}

	var filtrado []*Transaction
	for _, t := range *listaTransactions { // filtrado por todos los campos
		if tt.Id == t.Id && tt.CodigoTransaccion == t.CodigoTransaccion && tt.Moneda == t.Moneda && tt.Emisor == t.Emisor {
			filtrado = append(filtrado, &t)
		}
	}

	c.JSON(http.StatusOK, filtrado) // devovemos el array filtrado
}

func generarTransactions() *[]Transaction {
	transactions := []Transaction{
		{Id: 1, CodigoTransaccion: "abc123", Moneda: "peso", Monto: 100, Emisor: "Martín", Receptor: "Luisa"},
		{Id: 2, CodigoTransaccion: "abc134", Moneda: "dolar", Monto: 200, Emisor: "Marcos", Receptor: "Luisa"},
	}
	return &transactions
}
