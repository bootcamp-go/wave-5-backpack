package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre    string `json:"nombre"`
	Apellidos string `json:"apellidos"`
	Edad      int    `json:"edad"`
}

func main() {
	p := Persona{"Celeste", "Rodriguez", 23}
	miJson, _ := json.Marshal(p)

	fmt.Println(string(miJson))
	// fmt.Println(err)
}
