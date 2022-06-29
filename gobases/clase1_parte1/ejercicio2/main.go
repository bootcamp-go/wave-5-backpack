package main

import "fmt"

var (
	temperatura = 16
	humedad     = 62
	presion     = 1018.1 /*float*/
)

func main() {

	var slice = []int{1, 2, 3}
	var slice2 = []int{4, 5, 6}

	slice = append(slice, slice2...)

	fmt.Print("Temperatura: ", temperatura, " Humedad: ", humedad, " Presión: ", slice, "\n")

}
