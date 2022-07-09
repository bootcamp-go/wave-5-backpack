package main

import (
	"fmt"
)

func main() {

	age := 28
	employee := true
	workingSeniority := 22
	salary := 120000
	if age > 22 && employee == true && workingSeniority >= 1 && salary > 100000 {
		fmt.Println("Otorgado el prestámo")
	} else if age < 22 && employee == true && workingSeniority >= 1 && salary > 100000 {
		fmt.Println("No  cumples con la edad requerida para un prestámo")
	} else if employee == false && workingSeniority < 1 || salary < 100000 {
		fmt.Println("No cumples con las condiciones de empleo  y salario requerida")
	}

}
