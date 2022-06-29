package main

import "fmt"

func ex6() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	name := "Benjamin"
	fmt.Println("Benjamin tiene ", employees[name])
	fmt.Print("Los empleados con mas de 21 años son: ")
	for key, value := range employees {
		if value > 21 {
			fmt.Print(key, " ")
		}
	}
	fmt.Println(" ")
	delete(employees, "Pedro")
	fmt.Println(employees)
}
