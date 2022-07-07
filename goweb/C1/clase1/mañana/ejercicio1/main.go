package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*Ejercicio 1 - Estructura un JSON

Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no),
fecha de creación.
Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor
(string), fecha de transacción.
	1. Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el
	tema elegido, ej: products.json.
	2. Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o
	transacciones con todas sus variantes.
*/

type Transaction struct {
	Id                int64
	CodigoTransaccion string
	Moneda            string
	Monto             float64
	Emisor            string
	Receptor          string
}

func main() {
	transactions := []Transaction{
		{Id: 1, CodigoTransaccion: "abc123", Moneda: "peso", Monto: 100, Emisor: "Martín", Receptor: "Luisa"},
		{Id: 2, CodigoTransaccion: "abc134", Moneda: "dolar", Monto: 200, Emisor: "Marcos", Receptor: "Luisa"},
	}

	jsonTrasnsactions, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonTrasnsactions))
	os.WriteFile("./go-web/transactions.json", jsonTrasnsactions, 0644)
}
