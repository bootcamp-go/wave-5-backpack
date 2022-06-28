package main

import "fmt"

func main() {

	var (
		nombre    = "Pablo"
		direccion = "Sunchales"
	)
	var temp float32 = 13
	var moisture int8 = 64
	var presion float32 = 1018

	fmt.Println("-----------------------------------")
	fmt.Println("Hello Mundo! soy " + nombre + " y vivo en " + direccion)
	fmt.Println("En este momento la temperatura es de", temp, "°C")
	fmt.Println("La humedad es del", moisture, "%")
	fmt.Println("y la presión es de", presion, "mb")
	fmt.Println("-----------------------------------")
}
