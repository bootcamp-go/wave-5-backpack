package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	var (
		empleado string
		conteo   int
	)
	fmt.Print("Consulta empleado: ")
	fmt.Scanln(&empleado)
	fmt.Printf("%s tiene %d\n", empleado, employees[empleado])
	for _, e := range employees {
		if e > 21 {
			conteo++
		}
	}
	fmt.Printf("La empresa cuenta con %d empleados mayores de 21 anos\n", conteo)
	employees["Federico"] = 25
	fmt.Print(employees)
	delete(employees, "Pedro")
	fmt.Print(employees)
}
