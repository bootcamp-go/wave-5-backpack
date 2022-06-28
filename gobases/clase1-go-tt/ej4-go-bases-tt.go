/*Ejercicio 4 - Qué edad tiene...*/

package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Printf("\n|| Que edad tiene ||\n LISTA => ")
	fmt.Println(employees)

	// Saber cuántos de sus empleados son mayores de 21 años.
	fmt.Printf("\n > Empleados mayores a 21: \n")
	for name, age := range employees {
		if age >= 21 {
			fmt.Println("	Name: ", name, " Age:", age)
		}
	}

	// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	fmt.Printf("\n > Agregando al empleado 'Federico': \n")
	employees["Federico"] = 25
	fmt.Println(employees)

	// Eliminar a Pedro del mapa.
	fmt.Printf("\n > Eliminando al empleado 'Pedro': \n")
	delete(employees, "Pedro")
	fmt.Println(employees)
}
