// Ejercicio 4 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

package main

import (
	"fmt"
)

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	var empleadosMayores int
	for _, v := range employees {
		if v > 21 {
			empleadosMayores++
		}
	}

	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println("La edad de benjamin es:", employees["Benjamin"])
	fmt.Println("Cantidad de empleados mayores a 21 anios: ", empleadosMayores)
	fmt.Println("Imprimo lista: ", employees)

}
