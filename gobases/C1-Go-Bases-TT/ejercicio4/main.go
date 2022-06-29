package main

import "fmt"

// Ejercicio 4 - Qué edad tiene...
// Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin.

// var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

// Por otro lado también es necesario:
// Saber cuántos de sus empleados son mayores de 21 años.
// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
// Eliminar a Pedro del mapa.

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayoresDe21 = 0

	// Edad de Benjamin
	fmt.Println("Edad de Benjamin: ", employees["Benjamin"])

	// Cuántos son mayores de 21 años
	for _, e := range employees {
		if e > 21 {
			mayoresDe21++
		}
	}
	fmt.Println("Empleados mayores de 21 años: ", mayoresDe21)

	// Se agrega el empleado Federico de 25 años
	employees["Federico"] = 25

	// Se elimina el empleado Pedro
	delete(employees, "Pedro")

	// Se imprime la lista de empleados actualizados
	fmt.Println("Los empleados actualizados son: ")
	for k, e := range employees {
		fmt.Println(k, " => ", e)
	}
}
