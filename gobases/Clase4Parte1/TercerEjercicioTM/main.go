package main

import (
	"fmt"
)

func main() {
	salary := 15000

	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v\n", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
