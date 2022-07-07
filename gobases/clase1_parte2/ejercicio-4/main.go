package main

import "fmt"

func main() {
	//1. Edad de Benjamin
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad de Benjamin:", employees["Benjamin"])

	//2.
	// - Saber cuántos de sus empleados son mayores de 21 años.
	contador := 0
	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	fmt.Println("Cantidad de empleados mayores de 21 años:", contador)

	// - Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	employees["Federico"] = 25
	fmt.Println("Edad de Federico:", employees["Federico"])

	// - Eliminar a Pedro del mapa.
	delete(employees, "Pedro")
	fmt.Printf("Edad de Pedro %T:", employees["Pedro"])
}
