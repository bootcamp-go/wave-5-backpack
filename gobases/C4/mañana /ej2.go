package main

import "fmt"

func verificarImpuestos(salario int) (string, error) {
	if salario < 150000 {
		return "", &statusError{
			status: 400,
			msg:    "Error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else {
		return "Debe pagar impuestos", nil
	}

}

func main() {
	var salary int = 1540000

	status, err := verificarImpuestos(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(status)

}
