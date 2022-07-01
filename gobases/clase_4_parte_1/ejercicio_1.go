package main

import "fmt"

const MINIMO_IMPONIBLE = 150000

type imponibleError struct{}

func (imponibleError) Error() string {
	return "error: el salario ingresado no alcanza el minimo imponible"
}

func imponible(salary int) (string, error) {
	if salary < MINIMO_IMPONIBLE {
		return "", &imponibleError{}
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
