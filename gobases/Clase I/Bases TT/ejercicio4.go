package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es: %v \n", employees["Benjamin"])

	mayorDeEdad := 0
	for _, age := range employees {
		if age > 21 {
			mayorDeEdad += 1
		}
	}
	// Se imprime la cantidad empleados mayores a 21
	fmt.Printf("Hay %v mayores de 21 \n", mayorDeEdad)

	// Se agrega Federico y se imprime empleados
	employees["Federico"] = 25
	fmt.Println(employees)

	// Se elimina Pedro y se imprime empleados
	delete(employees, "Pedro")
	fmt.Println(employees)
}
