package main

import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}
type errorSalary struct {
	status int
	msg    string
}

func (e *errorSalary) Error() string {
	return fmt.Sprint("error: ", e.msg)
}
func main() {
	var salary int = 1252400

	err := testSalary(salary)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Debe pagar impuestos")
	}
}

func testSalary(salary int) error {
	if salary < 15000 {
		var statusCode int = 500
		return &errorSalary{
			status: statusCode,
			msg:    "el salario ingresado no alcanza el mínimo imponible",
		}
	}

	return nil
}
