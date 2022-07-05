/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Calcular promedio
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A school needs to calculate the average (per pupil) of its grades.
		It is requested to generate a function in which it can be passed
		N integers and return the average and an error if one of the numbers
		entered is negative.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	FUNCTION : promedio
func promedio(calificaciones ...float64) float64 {
	var length float64 = float64(len(calificaciones))
	var resultado_suma float64
	for _, calificacion := range calificaciones {
		resultado_suma += calificacion
		if calificacion < 0 {
			return -1
		}
	}
	return (resultado_suma / length)
}

//	MAIN PROGRAM
func main() {

	resultado := promedio(8, 7, 8, 4, 6, 7)
	fmt.Println("\n|| Calcular el promedio ||")

	if resultado < 0 {
		fmt.Println("Error. Un valor de la calificaciones es *negativo*")
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}
}
