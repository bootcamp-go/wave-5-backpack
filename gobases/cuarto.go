package main

import "fmt"

func main() {
	var count int = 0
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Printf("El empleado Benjamin tiene %d años de edad\n", employees["Benjamin"])

	for _, v := range employees {
		if v > 21 {
			count++
		}
	}
	fmt.Printf("Hay %d empleado(s) que tienen mas de 21 años\n", count)

	employees["Federico"] = 25
	fmt.Println("Se agrega al empleado Federico:", employees)

	delete(employees, "Pedro")
	fmt.Println("Se elimina al empleado Pedro:", employees)
}
