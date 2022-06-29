package main

import "fmt"

func queEdadTiene() {
	//ejercicio 4
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	count := 0

	for key := range employees {
		if employees[key] > 21 {
			count++
		}
	}

	fmt.Println("Empleados mayores de 21 años: ", count)

	employees["Federico"] = 25

	delete(employees, "Pedro")

}
