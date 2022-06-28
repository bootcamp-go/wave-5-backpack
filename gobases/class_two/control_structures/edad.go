package main

import "fmt"

func main() {
	var amountOfEmployeesMoreThan23 int = 0
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Darío":    44,
		"Pedro":    30,
	}

	for key, element := range employees {
		if key == "Benjamin" {
			fmt.Println(key, "has", element, "years")
		}
		if element > 21 {
			amountOfEmployeesMoreThan23++
		}
	}

	fmt.Println("There are", amountOfEmployeesMoreThan23, " that are older than 23")
	employees["Federico"] = 25
	delete(employees, "Pedro")

	fmt.Println(employees)
}
