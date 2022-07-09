package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println(mayores21(employees))

	employees["Federico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

}

func mayores21(e map[string]int) int {
	var employeeCounter int

	for _, v := range e {
		if v > 21 {
			employeeCounter++
		}
	}

	return employeeCounter
}
