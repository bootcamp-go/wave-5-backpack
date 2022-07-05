package main

import "fmt"

//Ejercicio 3 - Impuestos de salario

func validateSalary(salary int) error {
	if salary < 150000 {
		e := fmt.Errorf("error, el salario ingresado no alcanza el mínimo imponible")
		return e
	}
	return fmt.Errorf("Usted debe pagar impuestos %d", salary)

}

func main() {
	var salary int = 400000
	e := validateSalary(salary)
	fmt.Printf("Su estatus es: %v\n", e)
}
