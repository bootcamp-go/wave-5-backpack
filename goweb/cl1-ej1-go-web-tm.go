package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	/*
			transacciones:id,códigodetransacción(alfanumérico),moneda,monto,emisor
		(string), receptor (string), fecha de transacción
	*/
	type transaccion struct {
		Id                int     `json:"id"`
		CodigoTransaccion string  `json:"codigo de transaccion"`
		Moneda            string  `json:"moneda"`
		Monto             float64 `json:"monto"`
		Emisor            string  `json:"emisor"`
		Receptor          string  `json:"receptor"`
		Fecha             string  `json:"fecha de transaccion"`
	}

	t := transaccion{
		Id:                1,
		CodigoTransaccion: "1BFLSI",
		Moneda:            "EUR",
		Monto:             500.00,
		Emisor:            "Lalo",
		Receptor:          "Paco",
		Fecha:             "07/08/2020",
	}

	jsonData, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
