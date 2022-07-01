package main

import (
	"fmt"
)

func checkSalary(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

func main() {
	salary := 146000

	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
	}
}
