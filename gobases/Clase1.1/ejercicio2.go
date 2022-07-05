package main

import (
	"fmt"
)

func main() {
	var (
		temperature int     = 19
		humidity    float32 = 69.0
		pressure    float32 = 30.27
	)
	fmt.Printf("La temperatura es: %v \n", temperature)
	fmt.Printf("La humedad es: %v \n", humidity)
	fmt.Printf("La presion es: %v \n", pressure)

}
