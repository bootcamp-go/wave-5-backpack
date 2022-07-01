package main

import (
	"errors"
	"fmt"
	"os"
)

type errorSalary struct {
	status int
	msg    string
}

func main() {
	var salary int = 1000

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
		return errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}

	return nil
}
