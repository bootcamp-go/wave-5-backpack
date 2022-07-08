package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Maria": 30}
	count := 0
	fmt.Printf("La edad de Benjamin es: %d\n", employees["Benjamin"])

	for _, value := range employees {
		if value > 21 {
			count++
		}
	}

	fmt.Printf("La cantidad de empleados mayores a 21 es: %d\n", count)

	if _, value := employees["Pedro"]; !value {
		fmt.Printf("Pedro no está en la lista\n")
	} else {
		delete(employees, "Pedro")
		fmt.Printf("Pedro ha sido eliminado de la lista\n")
	}
}