/*
	Ejercicio 2 - Impuestos de salario #2

	Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, 
	en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
package main

import (
	"fmt"
	"errors"
)

func main() {
	var salary int = 20000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}