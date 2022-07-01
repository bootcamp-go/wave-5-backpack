package main

import (
	"errors"
	"fmt"
)

func errorValidationInt(salary int) (string, error) {
	if salary < 150000 {
		return "", errors.New("error: el salario ingresado no alcanza el mínimo disponible")
	}
	return "Debe pagar impuesto", nil
}

func main() {

	salary, err := errorValidationInt(160000)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(salary)

}
