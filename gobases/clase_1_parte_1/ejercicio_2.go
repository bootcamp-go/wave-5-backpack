package main

import "fmt"

func main(){
	var ( 
		temperatura int = 12
		humedad float32 = 62.2
		presion int = 1026
	)

	fmt.Println("Temperatura", temperatura, "°C")
	fmt.Println("% Humedad", humedad)
	fmt.Println("Presion", presion)
}