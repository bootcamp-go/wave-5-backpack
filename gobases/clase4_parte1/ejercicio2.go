package main

import (
	"errors"
	"fmt"
)

func checkTaxes(salary int) error {
	if salary < 150000 {
		return errors.New("Error: El salario ingresado no alcanza el minimo imponible")
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
