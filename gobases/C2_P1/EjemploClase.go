package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name     string `json:"Nombre"`
	LastName string `json:"Apellido"`
	DNI      int    `json:"Número_de_Identidad"`
	Address  string `json:"Dirección"`
}

func main() {
	p := Person{"Jose", "Riverón", 94061643089, "Av Circuito Merlot 2081"}
	jsonResult, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	r := reflect.TypeOf(p)

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		fmt.Println(field)
		tag := field.Tag.Get("json")
		fmt.Println(tag)
	}
	fmt.Println(r.Kind())
	fmt.Println(string(jsonResult))
}
