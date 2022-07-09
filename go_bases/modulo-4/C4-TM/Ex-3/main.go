package main

import (
	"fmt"
)

func main() {
	/*
		Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
		para que el mensaje de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible
		(el mensaje mostrado por consola deberá decir:
		“error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
	*/

	var salary int

	status, err := salaryCheck(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(status, "-Debe pagar impuesto 🔪")
}

func salaryCheck(s int) (int, error) {
	if s < 150000 {
		return 500, fmt.Errorf("❗️error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", s)
	}
	return 200, nil
}
