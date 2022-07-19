package main

import "fmt"

func main() {
	var temp float64
	var hum int
	var presion float64

	temp = 15.4
	hum = 65
	presion = 1014.3

	fmt.Printf("Temperatura actual: %v\n", temp)
	fmt.Printf("Humedad actual: %d\n", hum)
	fmt.Printf("Presión actual: %v\n", presion)

}
