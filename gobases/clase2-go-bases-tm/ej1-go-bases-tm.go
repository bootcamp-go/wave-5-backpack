/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #1:  Impuestos de salario
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A chocolate company needs to calculate the tax of its employees at the
		time of depositing the salary, to fulfill the objective it is necessary
		to create a function that returns the tax of a salary.
 		Taking into account that if the person earns more than $50,000, 17% will
		be deducted from the and if he/she earns more than $150,000, 10% will be
		deducted in addition.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	FUNCTION : impuesto
func impuesto(salarioEmpleado float64) float64 {
	var resultado float64

	if salarioEmpleado > 50000 && salarioEmpleado < 150000 {
		// Si posee mayor a 50,000 && menor a 150,000
		resultado = salarioEmpleado * .17
	} else if salarioEmpleado >= 150000 {
		// Mayor a 150,000
		resultado = salarioEmpleado * .1
	}
	return resultado
}

//	MAIN PROGRAM
func main() {
	//	Datos del empleado
	salario := 75000.0

	fmt.Println("\n|| Impuestos del salario ||")
	fmt.Printf("\nEl salario es %.2f \nTotal despues del impuesto: %.2f\n",
		salario, impuesto(salario))
}
