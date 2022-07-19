package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad de Benjamin:", employees["Benjamin"])

	contador := 0
	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	fmt.Println("Empleados mayores de 21 años:", contador)

	employees["Federico"] = 25
	fmt.Println("Edad de Federico:", employees["Federico"])

	delete(employees, "Pedro")
	fmt.Printf("Edad de Pedro %T:", employees["Pedro"])
}
