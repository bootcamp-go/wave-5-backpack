package main

import "fmt"

func main() {
	/* 
		var 1nombre string
		var apellido string
		var int edad
		1apellido := 6
		var licencia_de_conducir = true
		var estatura de la persona int
		cantidadDeHijos := 2
	*/
	
	var nombre string
	var apellido string
	var edad int
	edad = 34
	nombre = "Diana"
	apellido = "Quinteros"
	var licencia_de_conducir = true
	var estaturaDeLaPersona float64 = 1.66
	cantidadDeHijos := 2

	fmt.Println("nombre: ", nombre)
	fmt.Println("apellido: ", apellido)
	fmt.Println("edad: ", edad)
	fmt.Println("licencia de conducir: ", licencia_de_conducir)
	fmt.Println("estatura de la persona: ", estaturaDeLaPersona)
	fmt.Println("Cantidad de Hijos ", cantidadDeHijos)
}