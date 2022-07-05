/*---------------------------------------------------------------------------------*

     Assignment:	Bonus Track:  Impuestos de salario #4
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		Let's make our program a little more complex and useful.
		1. Develop the necessary functions to allow the company to calculate:
			(a) Monthly salary of a worker according to the number of hours worked.
			- The function will receive the hours worked in the month and the value
			  of the hour as an argument.
			- The function must return more than one value (calculated salary and
			  error).
			- In the event that the monthly salary is equal to or greater
			  than $150,000, a 10% tax deduction must be made.
			- If the number of monthly hours entered is less than 80 or
			  a negative number, the function must return an error. It should
			  indicate "error: the worker cannot have worked less than
			  80 hours per month".
			b) Calculate the half Christmas bonus corresponding to the worker
			- Bonus calculation formula:
			  [best salary of the semester] / 12 * [months worked in the semester].
			- The function that performs the calculation must return more than
			  one value, including an error in case it is entered. an error
			  if a negative number is entered.
		2. Develop the necessary code to fulfill the required functionalities,
			  using. "errors.New() ", "fmt.Errorf() " and " errors.Unwrap() ".  Do not forget to perform the validations of the error returns in your " main() " function
	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

package main

import (
	"errors"
	"fmt"
)

type errorAguinaldo struct{}

func (e errorAguinaldo) Error(cantidad float64) error {
	return fmt.Errorf("error. La cantidad de salario ingresada no debe ser negativo: %.2f", cantidad)
}

func salarioMensual(hora float64, valorHora float64) (float64, error) {
	salarioMensual := hora * valorHora

	// Si horas <= 80 entonces retorna un *error*
	if hora <= 80 {
		return 0, errors.New("error. El trabajador no puede haber trabajado menos de 80 hrs mensuales.")
	}

	// SI salario-mensual >= 150,000 se le (resta) 10%
	if salarioMensual >= 150000 {
		salarioMensual -= salarioMensual * 0.10
	}

	return salarioMensual, nil
}

func medioAguinaldo(salarioSemestre ...float64) (float64, error) {
	var totalMedioAguinaldo float64
	var mejorPagado float64
	var mesesTrabajados int

	for _, cantidad := range salarioSemestre {

		mesesTrabajados++

		if cantidad <= 0 {
			// Si el no. es negativo -> return error
			err := errorAguinaldo{}.Error(cantidad)
			return 0, fmt.Errorf("%w \n", err)
		} else if cantidad > mejorPagado {
			// La cantidad mayor se guarda
			mejorPagado = cantidad
		}
	}

	totalMedioAguinaldo = mejorPagado / 12 * float64(mesesTrabajados)
	return totalMedioAguinaldo, nil
}

func main() {
	fmt.Println("\n\t|| Impuestos de Salario #4 ||")

	totalSalario, err := salarioMensual(50, 1200)
	if err != nil {
		fmt.Println("\n ", err)
	} else {
		fmt.Printf("> El salario mensual es %.3f\n", totalSalario)
	}

	totalAguinaldo, err := medioAguinaldo(100, -50, 300, -5, 500)
	if err != nil {
		msgErr := errors.Unwrap(err)
		fmt.Println(" ", msgErr)
	} else {
		fmt.Printf("> El aguinaldo del semestre es %.3f \n", totalAguinaldo)
	}
}
