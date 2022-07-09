package main

import "fmt"

var nombre string = "José Luis"
var apellido string
var edad int = 28
var licenciaDeConducir bool = true
var estaturaDeLaPersona float64 = 1.74

func main() {
	apellido = "Riverón"
	cantidadDeHijos := 0
	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Apellido: %s\n", apellido)
	fmt.Printf("Edad: %d\n", edad)
	fmt.Printf("Licencia de Conducir: %v\n", licenciaDeConducir)
	fmt.Printf("Estatura: %v\n", estaturaDeLaPersona)
	fmt.Printf("Cantidad de Hijos: %d\n", cantidadDeHijos)
}
