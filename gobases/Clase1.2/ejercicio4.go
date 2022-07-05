package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int = 0

	fmt.Printf("Edad de benja: %v\n", employees["Benjamin"])

	for i := range employees {
		if employees[i] > 21 {
			mayores++
		}
	}

	fmt.Printf("Numero de empleados mayores de 21: %v\n", mayores)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
