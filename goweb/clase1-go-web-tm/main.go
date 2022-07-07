/*----------------------------------------------------------------------------------------*

     Assignment:	Ejercicio ##:  Go Web
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		Exercise#1
			1. Inside the goweb folder create a file [theme].json, the name has to
			be the chosen theme, ex: products.json.
			2. Inside it write a JSON that allows to have an array of products,
			users or transactions with all its variants.
		Exercise#2
			1. Create a file called main.go inside the go-web folder.
			2. Create a web server with Gin that responds you a JSON that has a
			   key "message" and says Hello followed by your name.
			3. Paste it to the endpoint to corroborate that the answer is correct.
		Exercise#3
			1. Inside the "main.go", create a structure according to the thematic
			   with the corresponding fields.
			2. Generate an endpoint whose path is/thematic (plural), e.g. "/products".
			3. Generate a handler for the endpoint called "GetAll".
			4. Creates an ices of the structure, then returns it through our endpoint.

	© Mercado Libre - IT Bootcamp 2022

*----------------------------------------------------------------------------------------*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaccion struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo de transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	Fecha             string  `json:"fecha de transaccion"`
}

func GetAll(w http.ResponseWriter, req *http.Request) {

	t := transaccion{
		Id:                1,
		CodigoTransaccion: "1BFLSI",
		Moneda:            "EUR",
		Monto:             500.00,
		Emisor:            "Lalo",
		Receptor:          "Paco",
		Fecha:             "2020-08-07",
	}

	jsonData, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, string(jsonData))
}

func main() {

	//	Ejercicio 2 - TM
	// Crea un router gon GIN [Framework]
	router := gin.Default()

	// Captura la solicitud GET "/message""
	router.GET("/message", func(c *gin.Context) {
		c.JSON(401, gin.H{"message": "Hola Israel! 👋"})
	})

	// Corremos nuestro servidor sobre el puerto 808
	router.Run(":8080")

	// Ejercicio 3 - TT
	// Captura la solicitud GET "/message""
	http.HandleFunc("/transacciones", GetAll)

	// Corremos nuestro servidor sobre el puerto 8080
	http.ListenAndServe(":8080", nil)
}
