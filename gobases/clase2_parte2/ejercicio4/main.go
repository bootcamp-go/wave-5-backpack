package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	mayores21 := 0
	for _, edad := range employees {
		if edad > 21 {
			mayores21++
		}
	}

	fmt.Println("La edad de Benjamin es :", employees["Benjamin"])
	fmt.Println("La cantidad de empleados mayores a 21 es :", mayores21)

	fmt.Println("Mapa antes de los cambios :", employees)
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Println("Mapa despues de los cambios :", employees)
}
