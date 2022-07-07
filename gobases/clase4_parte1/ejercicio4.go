/*
	Bonus Track -  Impuestos de salario #4
	Vamos a hacer que nuestro programa sea un poco más complejo.

	1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
		a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
			- La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
			- Dicha función deberá retornar más de un valor (salario calculado y error).
			- En caso de que el salario mensual sea igual o superior a $150.000, se le deberá 
			  descontar el 10% en concepto de impuesto.
			- En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número 
			  negativo, la función debe devolver un error. El mismo deberá indicar “error: 
			  el trabajador no puede haber trabajado menos de 80 hs mensuales”.

		b) Calcular el medio aguinaldo correspondiente al trabajador 
			- Fórmula de cálculo de aguinaldo: 
				[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
			- La función que realice el cálculo deberá retornar más de un valor, incluyendo un error 
			  en caso de que se ingrese un número negativo.

	2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando 
	“errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de 
	los retornos de error en tu función “main()”.
*/
package main

import (
	"fmt"
	"errors"
	"os"
)

func calculaSalario(horas float64, valorHora float64) (float64, error) {
	if horas < 80.0 {
		err := fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales, y trabajó: %.1f hs", horas)
		return 0, err
	}

	salario := horas * valorHora
	if salario >= 150000.0 {
		salario = salario * 0.9
	}
	return salario, nil
}

func calculaAguinaldo(mejorSalario float64, mesesTrabajados float64) (float64, error) {
	if mejorSalario < 0 || mesesTrabajados < 0 {
		err := errors.New("No puede ingresar números negativos")
		return 0, err
	}

	aguinaldo := mejorSalario / 12.0 * mesesTrabajados
	return aguinaldo, nil
}

func main() {
	horas := 90.0
	valorHora := 10000.0
	mejorSalario := 1000000.0
	mesesTrabajados := 5.0

	salario, err := calculaSalario(horas, valorHora)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("El salario calculado es:", salario)

	aguinaldo, err := calculaAguinaldo(mejorSalario, mesesTrabajados)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("El aguinaldo calculado es: %.0f \n", aguinaldo)
}