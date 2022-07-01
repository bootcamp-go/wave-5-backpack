package main

import (
	"fmt"
)

func errorValidationInt(salary int) (string, error) {
	if salary < 150000 {
		return "", fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de %d", salary)
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
