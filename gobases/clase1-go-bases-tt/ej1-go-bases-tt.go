/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Letras de una palabra
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		1. Create an application that has a variable with the word and print
		   the number of letters it has.
		2. Then I printed each of the letters.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	MAIN PROGRAM
func main() {
	text := "Academia"
	fmt.Printf("\n|| Letras de una palabra || \n\nPalabra designada: %s\n", text)
	fmt.Printf("> Cantidad de letras: %d \n", len(text))
	for _, index := range text {
		fmt.Print(" ", string(index))
	}
}
