package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var cantidad int
	for _, element := range employees {
		if element > 21 {
			cantidad++
		}
	}
	fmt.Println("Empleados mayores a 20 años son:", cantidad)
	// agregar un nuevo empleado
	employees["Enrique"] = 32
	// nueva lista de empleados
	fmt.Println(employees)
	// eliminar empleado
	delete(employees, "Nahuel")
	fmt.Println(employees)
}
