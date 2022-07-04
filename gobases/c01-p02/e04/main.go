package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println(employees)

	// #1
	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	// #2
	contador := 0
	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	fmt.Println("La cantidad de empleados mayor a 21 años es:", contador)

	// #3
	employees["Federico"] = 25
	fmt.Println(employees)

	// #4
	delete(employees, "Pedro")
	fmt.Println(employees)

}
