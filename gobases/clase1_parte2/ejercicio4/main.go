package main

import "fmt"

func mayores21(m map[string]int) int {
	var count int
	for _, v := range m {
		if v > 21 {
			count++
		}
	}

	return count
}

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Printf("Edad de Benjamin: %d\n", employees["Benjamin"])

	fmt.Printf("Tiene %v empleados mayores de 21\n", mayores21(employees))

	employees["Federico"] = 25
	fmt.Println("Agrega un empleado")
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println("Elimina un empleado del mapa")
	fmt.Println(employees)
}
