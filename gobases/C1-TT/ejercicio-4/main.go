package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("La edad de Benjamin es: %d\n", employees["Benjamin"])

	employees["Federico"] = 25
	delete(employees, "Pedro")

	mayores := 0
	for key := range employees {
		if employees[key] > 21 {
			mayores++
		}
	}
	fmt.Printf("Empleados mayores de 21 años: %d\n", mayores)
}
