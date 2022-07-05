package main

import "fmt"

func main() {
	empleados := map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	buscarEmpleado := "Benjamin"
	fmt.Printf("la edad del empleado %s, es %d\n", buscarEmpleado, empleados[buscarEmpleado])

	counter := 0
	for _, edad := range empleados {
		if edad > 21 {
			counter++
		}
	}
	fmt.Printf("hay %d empleados mayores de 25\n", counter)

	empleados["Federico"] = 25
	delete(empleados, "Pedro")

	fmt.Println(empleados)
}
