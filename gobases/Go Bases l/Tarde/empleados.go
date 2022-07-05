package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
var nombre string

func main() {

	menu := "Ingrese el nombre del empleado"

	fmt.Println(menu)

	fmt.Scanln(&nombre)

	fmt.Println(nombre, "tiene:", employees[nombre], "años")

	count := 0
	for _, edad := range employees {
		if edad > 21 {
			count++
		}
	}

	fmt.Println("Hay ", count, " empleados mayores de 21 años")

	employees["Federico"] = 25

	delete(employees, "Pedro")
	fmt.Println(employees)

}
