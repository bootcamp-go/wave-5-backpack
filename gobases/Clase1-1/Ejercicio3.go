package main

import "fmt"

func main() {
    var nombre string
	var apellido string
	var edad int 
	apellido = "tapia"
	var licenciaDeConducir = true
	var estaturaDeLaPersona float64
	cantidadDeHijos := 2

	nombre = "luis"
	edad = 25
	estaturaDeLaPersona = 1.67


	fmt.Printf("%v\n",nombre)
	fmt.Printf("%v\n",apellido)
	fmt.Printf("%v\n",edad)
	fmt.Printf("%v\n",licenciaDeConducir)
	fmt.Printf("%v\n",estaturaDeLaPersona)
	fmt.Printf("%v\n",cantidadDeHijos)
}