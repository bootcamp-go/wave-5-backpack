package main

import (
	"errors"
	"fmt"
)

func validateSalary(salary int) error {
	if salary < 150000 {
		return errors.New("error, el salario ingresado no alcanza el mínimo imponible")
	}
	return errors.New("Debe pagar impuestos")
}

func main() {
	var salary int = 40000
	e := validateSalary(salary)
	fmt.Printf("Su estatus es: %v\n", e)
}
