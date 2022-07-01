package main

import (
	"fmt"
)

type myCustomError struct {
	status int
	msg    string
}

func (err *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", err.status, err.msg)
}

func errorValidationInt(salary int) (string, error) {
	if salary < 150000 {
		return "", &myCustomError{
			status: 500,
			msg:    "error: el salario ingresado no alcanza el mínimo disponible",
		}
	}
	return "Debe pagar impuesto", nil
}

func main() {

	salary, err := errorValidationInt(140000)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(salary)

}
