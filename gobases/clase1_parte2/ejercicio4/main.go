package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var nombre = "Benjamin"

	var mayor = 0

	for _, v := range employees {

		if v > 21 {
			mayor++
		}

	}

	fmt.Printf("Edad Benjamin: %v \n", employees[nombre])

	fmt.Printf("cantidad mayor: %v \n", mayor)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}
