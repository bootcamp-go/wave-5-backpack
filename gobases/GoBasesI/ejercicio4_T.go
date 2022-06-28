package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Imprimiendo la edad de Benjamin
	fmt.Printf("La edad de Benjamin es %d \n", employees["Benjamin"])

	fmt.Println("---------------------------------")

	// Quienes y cuantos son mayores de 21
	var cont int = 0
	for key, value := range employees {
		if value > 21 {
			cont += 1
			fmt.Printf("%s tiene %d de edad \n", key, value)
		}
	}
	fmt.Printf("Hay %d mayores de 21 \n", cont)

	fmt.Println("---------------------------------")
	// Agregando un nuevo empleado a la lista
	employees["Federico"] = 21

	// Eliminando a Pedro del map
	fmt.Println(employees)
	fmt.Println("---------------------------------")
	delete(employees, "Pedro")
}
