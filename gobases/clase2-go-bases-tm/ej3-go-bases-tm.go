/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #3:  Calcular salario
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		A seafaring company needs to calculate the salary of its employees
		based on the number of hours worked per month and the category.
			If they are category C, their salary is $1,000 per hour.
			If it is category B, its salary is $1,500 per hour plus %20
			of its monthly salary.
			If you are category A, your salary is $3,000 per hour plus %50 of
			your monthly salary.
		You are requested to generate a function that receives the number of
		minutes worked per month and the category as a parameter, and returns
		your salary.


	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

//	CONSTANTS
const (
	categoriaA = "Categoria A"
	categoriaB = "Categoria B"
	categoriaC = "Categoria C"
)

//	FUNCTIONS : categories A,B,C & calculoSalario
func ctg_A(cantidadTiempo float64) float64 {
	var salarioMensual float64
	salarioMensual = 3000 * (cantidadTiempo / 60)
	return salarioMensual + (salarioMensual * .5)
}

func ctg_B(cantidadTiempo float64) float64 {
	var salarioMensual float64
	salarioMensual = 1500 * (cantidadTiempo / 60)
	return salarioMensual + (salarioMensual * .2)
}

func ctg_C(cantidadTiempo float64) float64 {
	var salarioMensual float64
	salarioMensual = 1000 * (cantidadTiempo / 60)
	return salarioMensual
}

func calculoSalario(categoria string, cantidadTiempo float64) float64 {
	switch categoria {
	case categoriaA:
		return ctg_A(cantidadTiempo)
	case categoriaB:
		return ctg_B(cantidadTiempo)
	case categoriaC:
		return ctg_C(cantidadTiempo)
	}

	return 0
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Calcular del Salario ||")
	// Datos del empleado
	cantidadTiempo := 120.0 // min. / mes

	fmt.Printf("> Categoria del empleado: %s \n> Cantidad de tiempo: %.2f (min)",
		categoriaB, cantidadTiempo)
	fmt.Println("\nSalario total: ", calculoSalario(categoriaB, cantidadTiempo))
}
