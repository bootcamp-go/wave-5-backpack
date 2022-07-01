package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Edad de benjamin: ", employees["Benjamin"])

	mayoresA21 := 0
	for _, edad := range employees {
		if edad > 21 {
			mayoresA21++
		}
	}

	employees["Francisco"] = 25

	delete(employees, "Pedro")

	fmt.Println("Empleados mayores a 21 años: ", mayoresA21)

	fmt.Println("Map después de las operaciones: ", employees)
}
