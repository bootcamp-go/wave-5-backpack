package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	cont := 0

	for _, edad := range employees {
		if edad > 21 {
			cont = cont + 1
		}
	}
	fmt.Println(employees)
	fmt.Println("Los empleados mayores de 21 años son ", cont)

	employees["Federico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
