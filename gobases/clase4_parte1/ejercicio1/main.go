package main

import "fmt"

type customError struct {
}

func (c *customError) Error() string {
	return "error : el salario ingresado no alcanza el minimo imponible"
}

func validarSalario(salary int) error {
	if salary < 150000 {
		return &customError{}
	}
	return nil
}

func main() {
	salary := 100000

	err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
	}

}
