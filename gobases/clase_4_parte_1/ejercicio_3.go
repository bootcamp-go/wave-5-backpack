package main

import (
	"fmt"
)

const MINIMO_IMPONIBLE = 150000

type imponibleError struct {
}

func (imponibleError) Error() string {
	return "error: el mínimo imponible es de " + fmt.Sprint(MINIMO_IMPONIBLE) + " y el salario ingresado es de: %d"
}

func imponible(salary int) (string, error) {
	if salary < MINIMO_IMPONIBLE {
		return "", fmt.Errorf(imponibleError{}.Error(), salary)
	}
	return "Debe pagar Impuesto", nil
}

func printMessageOrError(msg string, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
}

func main() {
	salary := 100000
	printMessageOrError(imponible(salary))
	salary = 160000
	printMessageOrError(imponible(salary))
}
