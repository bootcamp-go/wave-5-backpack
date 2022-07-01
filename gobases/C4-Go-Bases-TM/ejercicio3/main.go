package main

import (
	"fmt"
)

// Ejercicio 3 - Impuestos de salario #3

// Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”,
// para que el mensaje de error reciba por parámetro el valor de “salary” indicando
// que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir:
// “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”,
// siendo [salary] el valor de tipo int pasado por parámetro).

func validarSalario(salario int) error {
	if salario < 150000 {
		return fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salario)
	}
	return nil
}

func main() {
	fmt.Println("Ejercicio 3 - Impuestos de salario #3")
	fmt.Println("")

	salary := 1500

	err := validarSalario(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Debe pagar impuesto")
}
