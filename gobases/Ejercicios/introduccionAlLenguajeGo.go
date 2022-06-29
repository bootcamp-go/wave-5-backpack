package main

import "fmt"

func imprimiTuNombre() {
	//ejercicio 1
	nombre := "Francisco"
	direccion := "Antonio Podio 1220"

	fmt.Print("Nombre: ", nombre, "\nDireccion: ", direccion)
}

func clima() {
	//ejercicio 2
	var temperatura int = 15
	var humedad int = 90
	var presion float32 = 1001.1

	fmt.Print("\nTemperatura: ", temperatura, "ºC", "\nHumedad: ", humedad, "%", "\nPresion: ", presion)
	fmt.Println()
}

/*
ejercicio 3
var nombre string
var apellido string
var edad int
apellido := "Monay"
var licencia_de_conducir bool = true
var estaturaDeLaPersona float32
cantidadDeHijos := 2
*/

/*
ejercicio 4
var apellido string = "Gomez"
var edad int = 35
boolean := false
var sueldo float32 = 45857.90
var nombre string = "Julián"
*/
