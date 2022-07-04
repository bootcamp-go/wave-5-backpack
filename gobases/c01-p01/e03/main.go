package main

import "fmt"

func main() {

	// var 1nombre string
	// Los nombres de variables no pueden iniciar con un número
	var nombre string
	nombre = "Abelardo"

	var apellido string // Correcto
	apellido = "Lugo"

	// var int edad
	// Primero se escribe el nombre de la variable y luego el tipo de dato
	var edad int
	edad = 30

	// 1apellido := 6
	// Los nombres de variables no pueden iniciar con un número
	_apellido := 6

	var licencia_de_conducir = true // Correcto

	// var estatura de la persona int
	// Los nombres no pueden tener espacios
	var estatura_de_la_persona int
	estatura_de_la_persona = 2

	cantidadDeHijos := 2 // Correcto

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Println(edad)
	fmt.Println(_apellido)
	fmt.Println(licencia_de_conducir)
	fmt.Println(estatura_de_la_persona)
	fmt.Println(cantidadDeHijos)

}
