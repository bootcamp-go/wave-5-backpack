/*Ejercicio 3 - A qué mes corresponde*/

package main

import "fmt"

func main() {
	meses := []string{
		"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre",
	}
	seleccion := 13

	fmt.Printf("\n|| A que corresponde ||\n")
	if seleccion > 0 && seleccion < 13 {
		fmt.Printf("El mes seleccionado %d es %s", 2, meses[2])
	} else {
		fmt.Printf("Selecciona un numero del 1 al 12")
	}
}
