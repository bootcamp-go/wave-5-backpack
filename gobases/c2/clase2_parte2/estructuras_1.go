package main

import "fmt"

type Persona struct {
	Nombre    string
	Apellidos string
	Edad      int
}

func main() {

	// var p0 Persona

	// fmt.Printf("%+v\n", p0)

	// p1 := Persona{"Nicolas", "Aldeco", 23}

	// fmt.Printf("%+v\n", p1)

	p2 := Persona{
		Nombre:    "Nicolas",
		Apellidos: "Aldeco",
	}

	// fmt.Printf("%+v\n", p2)

	fmt.Printf("p2.Nombre: %s\n", p2.Nombre)

	p2.Nombre = "Ariel"

	fmt.Printf("p2.Nombre: %s\n", p2.Nombre)

}
