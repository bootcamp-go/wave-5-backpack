package main

import (
	"fmt"
)

func checkImpuesto(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("El mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	} else {
		fmt.Println("Debe pagar impuestos")
		return nil
	}
}

func main() {
	salary := 140000
	err := checkImpuesto(salary)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	salary = 170000
	err = checkImpuesto(salary)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
