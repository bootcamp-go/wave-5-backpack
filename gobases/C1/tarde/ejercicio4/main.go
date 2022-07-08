package main

import "fmt"

/*Ejercicio 4 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa,
ayuda  a imprimir la edad de Benjamin.

Por otro lado también es necesario:
	1. Saber cuántos de sus empleados son mayores de 21 años.
	2. Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	3. Eliminar a Pedro del mapa.
*/

func main() {
	employes := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	// Edad de Benjamin
	fmt.Printf("Edad de Benjamin %d\n", employes["Benjamin"])

	// Empleados mayores de 21
	var cantMayores21 int
	for _, v := range employes {
		if v > 21 {
			cantMayores21++
		}
	}
	fmt.Printf("Cantidad de empleados mayores de 21 = %d\n", cantMayores21)

	// Agregar a Federico
	employes["Federico"] = 25
	fmt.Println(employes)

	// Eliminar a Pedro
	delete(employes, "Pedro")
	fmt.Println(employes)
}
