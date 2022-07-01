package main

import (
	"fmt"
)

func checkTaxes(salary int) error {
	if salary < 150000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150000 y el salario ingresado es de %d", salary)
	} else {
		fmt.Println("Debe pagar impuestos")
		return nil
	}

}

func main() {
	var salary int = 100000
	err := checkTaxes(salary)
	if err != nil {
		fmt.Println(err)
	}

	salary = 160000
	err = checkTaxes(salary)
	if err != nil {
		fmt.Println(err)
	}
}
