package main

import (
	"fmt"
)

/*Ejercicio 3 - Impuestos de salario #3

Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje
de error reciba por parámetro el valor de “salary” indicando que no alcanza el mínimo
imponible (el mensaje mostrado por consola deberá decir: “error: el mínimo imponible es
de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de
tipo int pasado por parámetro).*/

func main() {
	var salary int = 200000
	if salary < 150000 {
		err := fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto.")
	}
}
