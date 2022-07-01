package main

import (
	"errors"
	"fmt"
)

// Funcion que realiza verificación de salario

func salaryVerification(salary int) error {

	if salary < 150000 {
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	return nil
}

func main() {

	salary := 100000
	err := salaryVerification(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Se debe pagar impuesto")
	}
}
