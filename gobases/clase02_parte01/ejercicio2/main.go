package main

import "fmt"

var(
	temperatura int
	humedad float32
	presion float32
)

func main(){
	temperatura = 15
	humedad = 0.36
	presion = 1000

	fmt.Println("Temperatura: ", temperatura, "\nHumedad: ", humedad, "\nPresion: ", presion)
}