package main

import (
	"errors"
	"fmt"
)

func checkSalary(salary int) error {
	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

func main() {
	salary := 156000

	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
	}
}
