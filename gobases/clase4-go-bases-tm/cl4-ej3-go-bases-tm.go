/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  Impuestos de salario #3
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		Repeat the previous process, but now implementing " fmt.Errorf() ",
		so that the error message receives by parameter the value of " salary "
		indicating that it does not reach the taxable minimum (the message displayed
		by console should read: " error: the taxable minimum is 150,000 and
		the salary entered is: [salary] ", where [ salary ] is the value of
		type int passed by parameter.
	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

/* Ejercicio 3 - Impuestos de salario #3 */

package main

import (
	"fmt"
	"os"
)

func msgError(salary int) error {
	if salary < 150000 {
		err := fmt.Errorf("error:  el mínimo imponible es de 150.000 y el salario ingresado es de: %d\n😱😱😱", salary)
		return err
	}
	return nil
}

func main() {
	fmt.Println("\n\t|| Impuestos de Salario #3 ||")
	var salary int
	salary = 100000

	err := msgError(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuesto 🤑 : $", salary)
}
