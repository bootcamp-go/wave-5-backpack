package main

import (
	"fmt"
	"os"
)

func main() {
	var salary int = 12400

	err := testSalary(salary)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}

func testSalary(salary int) error {
	if salary < 15000 {
		return fmt.Errorf("rror: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary)
	}

	return nil
}
