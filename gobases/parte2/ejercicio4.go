package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Benjamin tiene %d años\n", employees["Benjamin"])

	count := 0
	for _, element := range employees {
		if element > 21 {
			count++
		}
	}

	fmt.Printf("Empleados mayores de 21 son: %d\n", count)

	delete(employees, "Pedro")

	fmt.Println(employees)

}
