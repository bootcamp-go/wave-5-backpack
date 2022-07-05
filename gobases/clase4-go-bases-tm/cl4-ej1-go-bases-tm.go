/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Impuestos de salario #1
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		1. In your "main" function, define a variable called "salary" and
		   assign it a value of type "int".
		2. Create a custom error with a struct that implements
		  "Error()" with the message "error: the salary entered does not reach the minimum taxable salary" and throw it if "salary" is less than 150,000. Otherwise, it prints by console the message "Must pay tax".

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"fmt"
	"os"
)

//	STRUCT : customError
type customError struct {
	cantidad int
	msg      string
}

//	FUNCTIONS
func (e *customError) Error() string {
	return fmt.Sprintf("%v\n", e.msg)
}

func customErrorTest(cantidad int) error {
	if cantidad < 150000 {
		return &customError{
			cantidad: cantidad,
			msg:      "ERROR: el salario ingresado no alcanza el mínimo impoible 😞",
		}
	}
	return nil
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Impuestos de Salario #1 ||")
	var salary int
	salary = 170000

	err := customErrorTest(salary)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("Debe pagar impuesto 🤑 : $%v\n", salary)
}
