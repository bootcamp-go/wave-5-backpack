package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var mayores int = 0

	fmt.Println(employees["Benjamin"])

	for _, edad := range employees {
		if edad > 21 {
			mayores++
		}
	}

	fmt.Println(mayores)
	fmt.Println(employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}
