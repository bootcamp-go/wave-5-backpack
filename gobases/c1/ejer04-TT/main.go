package main

import "fmt"

// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados.
// Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

//   var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])

	total := 0

	for _, edad := range employees {

		if edad > 21 {
			total++
		}
	}
	fmt.Println(total)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)

}
