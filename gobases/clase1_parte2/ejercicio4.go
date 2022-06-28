package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	edadBenjamin(employees)
	mayoresEdad(employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}

func edadBenjamin(employees map[string]int) {
	for key, value := range employees {
		if key == "Benjamin" {
			fmt.Printf("La edad de Benjamin es %d años\n", value)
		}
	}
}

func mayoresEdad(employees map[string]int) {
	mayores := 0
	for _, value := range employees {
		if value > 21 {
			mayores++
		}
	}
	fmt.Printf("%d Empleados son mayores de 21 años\n", mayores)
}
