package main

import "fmt"

var temperature int = 29
var pressure float64 = 1026.8
var humidity int = 28

func main() {
	fmt.Printf("La temperatura es de: %dºC\n", temperature)
	fmt.Printf("La presión es de: %vmb\n", pressure)
	fmt.Printf("La humedad relativa es de: %d\n", humidity)
}
