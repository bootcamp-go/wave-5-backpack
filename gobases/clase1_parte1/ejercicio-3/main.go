package main

import "fmt"

func main() {
	/* Punto #1 */

	/*
		var 1nombre string 				// Las variables deben de iniciar con una letra
		var apellido string				// Correcto
		var int edad 					// Primero va el nombre y luego el tipo
		1apellido := 6 					// Las variables deben de iniciar con una letra
		var licencia_de_conducir = true	// Correcto
		var estatura de la persona int 	// Los nombres no pueden tener espacios
		cantidadDeHijos := 2			// Correcto
	*/

	/* Punto #2 */

	var nombre string
	var apellido string
	var edad int
	apellido1 := 6
	var licencia_de_conducir = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2

	// El código de imprimir va a dar errores porque hay variables que no están
	// inicializadas.
	fmt.Print(nombre, apellido, edad, apellido1, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)

}
