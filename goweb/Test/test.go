package main

import (
	"fmt"
	"reflect"
)

type request struct {
	CodigoTransaccion int     `json:"codigo_de_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_de_transaccion"`
}
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

func main() {

	r := request{1, "COP", 23.2, "foo", "e3e", "hoy"}
	p := Product{1, "elkin", "niño", 2, 2.2}

	p2 := request{CodigoTransaccion: 0, Moneda: "USD", Receptor: "no se"}

	validateStruct(r)
	validateStruct(p)
	validateStruct(p2)
}

func validateStruct(s interface{}) {

	v := reflect.ValueOf(s) // applied over a var of type struct
	t := reflect.TypeOf(s)
	values := make([]interface{}, v.NumField())
	names := make([]string, t.NumField())
	var errMsg string
	for i := range names {
		names[i] = t.Field(i).Name
		values[i] = v.Field(i).Interface()
		if IsZero[values[i].type](values[i]) {

			errMsg += names[i]
		}

		fmt.Printf("%T %v \n", values[i], values[i])
	}

	fmt.Println(names)
	fmt.Println(values)
	fmt.Println(errMsg)
}
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
