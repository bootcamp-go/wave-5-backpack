package main

import (
	"fmt"
)

type statusError struct {
	msg string
}

func (e *statusError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func checkImpuesto(salary int) error {
	if salary < 150000 {
		return &statusError{
			msg: "El salario ingresado no alcanza el mínimo imponible.",
		}
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
