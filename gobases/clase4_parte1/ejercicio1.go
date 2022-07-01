package main

import (
	"fmt"
)

type customError struct {
	code int
	msg  string
}

func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.msg)
}

func checkTaxes(salary int) error {
	if salary < 150000 {
		return &customError{
			code: 123,
			msg:  "El salario ingresado no alcanza el minimo imponible",
		}
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
