package main

import (
	"fmt"
)

func main() {
	salary := 50000
	if salary < 150000 {
		err := fmt.Errorf("el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}
