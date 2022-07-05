/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Clima
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		1. Declare 3 variables specifying the type of data, as value
		   they must have the temperature, humidity and pressure of where you are.
		2. Prints the values of the variables in the console.
		3. What type of data would you assign to the variables?

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

// PACKAGE & LIBRARY
package main

import "fmt"

// Creating variables
var temperatura, humedad int
var presion float64

// MAIN PROGRAM
func main() {
	temperatura := 18
	humedad := 67
	presion := 1018.0

	fmt.Println("\nClima:")
	fmt.Println("\tTemperatura: ", temperatura, "º")
	fmt.Println("\tHumedad: ", humedad, "%")
	fmt.Println("\tPresion: ", presion, " mb")
}
