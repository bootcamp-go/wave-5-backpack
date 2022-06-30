/*
	Ejercicio 4 - Qué edad tiene...
	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. 
	Según el siguiente mapa, ayuda  a imprimir la edad de Benjamin. 

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	Por otro lado también es necesario: 
	Saber cuántos de sus empleados son mayores de 21 años.
	Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	Eliminar a Pedro del mapa.
*/

package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	edadBenjamin := employees["Benjamin"]

	fmt.Println("Edad de Benjamín: ", edadBenjamin)

	// Saber cuántos de sus empleados son mayores de 21 años.
	mayoresA21 := 0
	for _, age := range employees {
		if age > 21 {
			mayoresA21++
		}
	}
	fmt.Println("Los empleados mayores a 21 son: ", mayoresA21)

	// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	employees["Federico"] = 25
	
	// Eliminar a Pedro del mapa.
	delete(employees, "Pedro")

	for name, age := range employees {
		fmt.Println("Nombre:", name, "| Edad:", age)
	}
}