package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	empleadosMayores := 0

	fmt.Println(employees["Benjamin"])

	for _, edad := range employees {
		if edad > 21 {
			empleadosMayores++
		}
	}

	fmt.Printf("%d empleados son mayores de 21 \n", empleadosMayores)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)

}
