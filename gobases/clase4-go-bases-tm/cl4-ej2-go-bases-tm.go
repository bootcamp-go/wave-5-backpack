/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Impuestos de salario #2
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		Do the same as in the previous exercise but reformulate the code so that,
		instead of "Error()", "errors.New()" is implemented.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"errors"
	"fmt"
)

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Impuestos de Salario #2 ||")
	var salary int
	salary = 100000

	if salary < 150000 {
		fmt.Println(errors.New("ERROR: el salario ingresado no alcanza el mínimo impoible 😞"))
		return
	}
	fmt.Println("Debe pagar impuesto 🤑 : $", salary)
}
