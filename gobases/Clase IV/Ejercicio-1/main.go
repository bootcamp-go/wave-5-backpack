package main

import (
	"fmt"
)

// Estructura implementada para Error

type mycustomError struct {
	msg string
}

// Funcion que struct implementa el método Error()

func (m *mycustomError) Error() string {

	return fmt.Sprintf("%v", m.msg)

}

// Funcion que realiza verificación de salario

func salaryVerification(salary int) error {

	if salary < 150000 {
		return &mycustomError{
			msg: "error: el salario ingresado no alcanza el mínimo imponible",
		}
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
