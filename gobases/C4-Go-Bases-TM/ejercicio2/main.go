package main

import (
	"errors"
	"fmt"
)

// Ejercicio 2 - Impuestos de salario #2

// Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
// se implemente “errors.New()”.

func validarSalario(salario int) {
	if salario < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")
}

func main() {
	fmt.Println("Ejercicio 2 - Impuestos de salario #2")
	fmt.Println("")

	salary := 1500

	validarSalario(salary)
}
