package main

import (
	"fmt"
)

// Funcion que realiza verificación de salario

func salaryVerification(salary int) error {

	if salary < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de %d", salary)
	}
	return nil
}

func main() {

	salary := 160000
	err := salaryVerification(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se debe pagar impuesto")
	}
}
