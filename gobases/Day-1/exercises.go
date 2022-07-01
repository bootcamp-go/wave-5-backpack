package main

import "fmt"

func main() {
	var (
		nombre      string = "Santiago Rafael Salcedo Camacho"
		direccion   string = "Carrera 30 22 12"
		temperatura int    = 30
		humedad     string = "67%"
		presion     int    = 1015
	)
	fmt.Println("1. Nombre y Direccion: \n", nombre, "-", direccion)
	fmt.Println("2. Clima \n Temperatura:", temperatura, "\n Humedad:", humedad, "\n Presion:", presion)
}
