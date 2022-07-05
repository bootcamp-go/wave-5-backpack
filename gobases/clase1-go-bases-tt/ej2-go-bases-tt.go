/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Préstamo
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A bank wants to grant loans to its customers, but not everyone can access
		them. For this, it has certain rules to know to which client it can be
		granted. It only grants loans to clients who are over 22 years old, are
		employed and have been in their job for more than a year. Within the
		loans it grants, it will not charge interest to those whose salary is
		better than $100,000. It is necessary to make an application that has
		these variables and that prints a message according to each case.
		Tip: your code must be able to print at least 3 different messages.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

// STRUCT: Data
type Data struct {
	Name     string
	Edad     int
	Empleado bool
	Tiempo   int
}

//	MAIN PROGRAM
func main() {

	employee := Data{"Samuel", 25, true, 2}

	fmt.Printf("\n|| Préstamo ||")
	fmt.Println("\nPersona: ", employee.Name)

	//CONDITIONS
	if employee.Edad > 22 {
		fmt.Printf("\n>> Posee mayor a 22 años :: Tiene %d años\n", employee.Edad)
	} else {
		fmt.Printf("\n>> No posse mayor a 22 años ")
	}
	if employee.Empleado == true {
		fmt.Printf(">> Si es empleado\n")
	} else {
		fmt.Printf("\n>> No es empleado ")
	}
	if employee.Tiempo > 1 {
		fmt.Printf(">> Posee mayor a 1 año :: Tiene %d (años)\n", employee.Tiempo)
	} else {
		fmt.Printf("\n>> Tiene menor a un año ")
	}
}
