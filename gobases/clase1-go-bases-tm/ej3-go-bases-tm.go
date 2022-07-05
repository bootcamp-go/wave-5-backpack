/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  Declaración de variables
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		1. Detect which of these variables declared by the student are correct.
		2. Correct the incorrect ones.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

/*NOTE: The variables were corrected according to the activity without being used.*/

// PACKAGE
package main

// MAIN PROGRAM
func main() {
	var nombre string              // Cambio en el 1nombre -> nombre
	var apellido int               // Cambio en tipo de asignacion string -> int
	var edad int                   // Cambiar en el orden correcto
	apellido = 6                   // Cambio en el 1apellido -> apellido && no requiere := sino =
	licencia_de_conducir := true   // Utilizar :=
	var estatura_de_la_persona int // Cambiar * sin espacios *
	cantidadDeHijos := 2
}
