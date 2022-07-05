/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  A qué mes corresponde
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		Make an application that contains a variable with the number of the month.
		1. Depending on the number, print the corresponding month in text.
		2. Can you think if it can be solved in more than one way?
		   Which would you choose and why?

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	MAIN PROGRAM
func main() {

	meses := []string{ // Slice - String
		"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
	}
	seleccion := 13

	fmt.Printf("\n|| A que corresponde ||\n")
	if seleccion > 0 && seleccion < 13 { // Condition for print 'meses'
		fmt.Printf("El mes seleccionado %d es %s", 2, meses[2])
	} else { // Otherwise, it prints an error message.
		fmt.Printf("Selecciona un numero del 1 al 12")
	}
}
