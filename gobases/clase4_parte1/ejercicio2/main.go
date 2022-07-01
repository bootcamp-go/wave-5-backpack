package main

import (
	"errors"
	"fmt"
)

func validarSalario(salary int) error {
	if salary < 150000 {
		return errors.New("error : el salario ingresado no alcanza el minimo imponible")
	}
	return nil
}

func main() {
	salary := 149999

	err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
	}

}
