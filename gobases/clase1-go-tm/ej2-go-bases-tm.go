/* Ejercicio 2 - Clima */

package main

import "fmt"

var temperatura, humedad int
var presion float64

func main(){
	temperatura:= 18
	humedad:= 67
	presion:= 1018.0

	fmt.Println("\nClima:")
	fmt.Println("\tTemperatura: ",temperatura,"º")
	fmt.Println("\tHumedad: ",humedad,"%")
	fmt.Println("\tPresion: ",presion," mb")
}