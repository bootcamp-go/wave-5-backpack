package main

import "fmt"

func main() {
	var empleadosMayoresDe21 int
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("Edad de Benjamin: %v \n", employees["Benjamin"])
	for _, value := range employees {
		if value >= 21 {
			empleadosMayoresDe21++
		}
	}
	fmt.Printf("Empleados mayores de 21: %v \n", empleadosMayoresDe21)
	employees["Federico"] = 25
	delete(employees, "Pedro")
}
