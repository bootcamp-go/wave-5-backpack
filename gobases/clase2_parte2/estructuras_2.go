package main

import "fmt"

type Preferencias struct {
	Pelicula, Comida string
}

type Persona struct {
	Nombre    string
	Apellidos string
	Edad      int
	Gustos    Preferencias
}

func main() {

	// p1 := Persona{"Nicolas", "Tesone", 23, Preferencias{"Titanic", "Asado"}}

	// fmt.Printf("p1: %+v \n", p1)

	p2 := Persona{
		Nombre:    "John",
		Apellidos: "Connor",
		Gustos: Preferencias{
			Pelicula: "Terminator",
			Comida:   "Ravioles",
		},
	}

	fmt.Printf("p2: %+v \n", p2)

	// var p3 Persona

	// p3.Gustos.Comida = "Arepas"
	// fmt.Printf("p3: %+v \n", p3)
}
