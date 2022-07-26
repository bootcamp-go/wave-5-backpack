package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre    string `bd:"nombre"`
	Apellidos string `bd:"apellidos"`
	Edad      int    `bd:"edad"`
}

func main() {
	persona := Persona{"Nicolas", "Aldeco", 23}
	personaReflect := reflect.TypeOf(persona)

	fmt.Printf("Type: %v\n", personaReflect.Name())
	fmt.Printf("Kind: %v\n\n", personaReflect.Kind())

	for i := 0; i < personaReflect.NumField(); i++ {
		field := personaReflect.Field(i)
		fmt.Printf("Tag: %v\n", field.Tag.Get("bd"))
	}

}
