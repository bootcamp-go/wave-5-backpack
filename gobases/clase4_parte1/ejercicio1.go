package main

import "fmt"

//Impuesto de salario #1

type errorSalary struct {
	msg string
}

func (e *errorSalary) Error() string {
	return fmt.Sprintln("error: el salario ingresado no alcanza el mínimo imponible")
}

func validateSalary(salary int) error {
	if salary < 150000 {
		return &errorSalary{}
	}
	return nil
}

func main() {
	var salary int = 400000
	if validateSalary(salary) != nil {
		fmt.Println(validateSalary(salary))
		return
	}
	fmt.Println("Alerta debe pagar impuestos")
}
