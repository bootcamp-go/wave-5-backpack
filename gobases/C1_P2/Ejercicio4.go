package main

import (
	"fmt"
)

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Benjamin age:", employees["Benjamin"])
	employees["Federico"] = 25
	delete(employees, "Pedro")
	for key, element := range employees {
		if element > 21 {
			fmt.Println("Employee Name:", key, element, "years old")
		}
	}
	fmt.Println(employees)
}
