package main

import "fmt"

// var 1nombre string 				XX Las variables deben de iniciar con una letra
// var apellido string				OO Correcto
// var int edad 					XX Primero va el nombre y luego el tipo
// 1apellido := 6 					XX Las variables deben de iniciar con una letra
// var licencia_de_conducir = true	OO Correcto
// var estatura de la persona int 	XX Los nombres no pueden tener espacios
// cantidadDeHijos := 2				OO Correcto
func main() {

	var nombre string
	var apellido string
	var edad int
	apellido1 := 6
	var licencia_de_conducir = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2

	fmt.Print(nombre, apellido, edad, apellido1, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)

}
