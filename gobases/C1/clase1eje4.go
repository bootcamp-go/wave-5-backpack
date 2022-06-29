package main

import (
	"fmt"
)

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin ",employees["Benjamin"])

	x := 0
	for _, edad := range employees {
		if edad > 21{
			x++
		}
	}
	fmt.Println("Numero de empleados con edad mayor a 21: ",x)

	employees["Federico"]= 25

	delete(employees,"Pedro")

	fmt.Println(employees)
}