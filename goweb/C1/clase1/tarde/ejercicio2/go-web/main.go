package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 2 - Get one endpoint

Generar un nuevo endpoint que nos permita traer un solo resultado del array de la temática.
Utilizando path parameters el endpoint debería ser /temática/:id (recuerda que siempre tiene
que ser en plural la temática). Una vez recibido el id devuelve la posición correspondiente.
	1. Genera una nueva ruta.
	2. Genera un handler para la ruta creada.
	3. Dentro del handler busca el item que necesitas.
	4. Devuelve el item según el id.
Si no encontraste ningún elemento con ese id devolver como código de respuesta 404.
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

	// Ejercico N°2
	router.GET("/transactions/:id", TransactionHandler)
	router.Run(":8000") // Por defecto gin arranca el server on 8080
}

func TransactionHandler(c *gin.Context) {
	listaTransactions := generarTransactions()
	var transaction Transaction

	id, _ := strconv.Atoi(c.Param("id"))
	find := false
	for _, t := range listaTransactions {
		if t.Id == int64(id) {
			find = true
			transaction = *t
			break
		}
	}

	if !find {
		c.JSON(http.StatusNotFound, nil) // 404
	} else {
		c.JSON(http.StatusOK, transaction) // 200 => transaction
	}
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
	for _, t := range listaTransactions { // filtrado por todos los campos
		if tt.Id == t.Id && tt.CodigoTransaccion == t.CodigoTransaccion && tt.Moneda == t.Moneda && tt.Emisor == t.Emisor {
			filtrado = append(filtrado, t)
		}
	}

	c.JSON(http.StatusOK, filtrado) // devovemos el array filtrado
}

func generarTransactions() []*Transaction {
	transactions := []*Transaction{
		{Id: 1, CodigoTransaccion: "abc123", Moneda: "peso", Monto: 100, Emisor: "Martín", Receptor: "Luisa"},
		{Id: 2, CodigoTransaccion: "abc134", Moneda: "dolar", Monto: 200, Emisor: "Marcos", Receptor: "Luisa"},
	}
	return transactions
}
