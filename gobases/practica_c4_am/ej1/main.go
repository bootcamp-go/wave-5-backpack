package main

import (
	"fmt"
)

type MiError struct{}

func (this *MiError) Error() string {
	return "error: el salario no cumple requisito."
}

func verificaciones(salary int) error {
	if salary < 150000 {
		return &MiError{}
	}
	return nil
}

func main() {
	var salary int = 10000000
	err := verificaciones(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}

	salary = 15000
	err = verificaciones(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
