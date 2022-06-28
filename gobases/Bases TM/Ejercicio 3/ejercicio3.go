package main

import "fmt"

// var 1nombre string  no puede empezar con un número
var nombre string
var apellido string
// var int edad  está al revés
var edad int
//1apellido:=6 no se puede comenzar con un número y está definido como string

//var estatura de la persona int no puede tener espacios el nombre de la variable
var estaturaDeLaPersona int

//var licencia_de_conducir=true falto indicar que es boleano

func main (){

	cantidadDeHijos:=2
	licencia_de_conducir:= true
	apellido= "Brito"

	fmt.Println(nombre,apellido,edad,licencia_de_conducir,estaturaDeLaPersona,cantidadDeHijos)

}