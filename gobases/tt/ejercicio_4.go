package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	// Edad de Benjamin
	fmt.Println("Edad de Benjamin:",employees["Benjamin"])
	
	// Empleados mayores a 21 años
	var employees21 []string
	for k, v := range employees {
		if (v > 21) {
			employees21 = append(employees21, k)
		}
	}
	fmt.Println("Empleados mayores de 21 años:", employees21)

	// Update del map
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Nueva lista de empleados:", employees)	
}