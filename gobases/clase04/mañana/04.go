/*
Bonus Track -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo.
1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
	a. Salario mensual de un trabajador según la cantidad de horas trabajadas.
		- La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
		- Dicha función deberá retornar más de un valor (salario calculado y error).
		- En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
		- En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.

	b. Calcular el medio aguinaldo correspondiente al trabajador 
		- Fórmula de cálculo de aguinaldo: [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
		- La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.

*/

package main

import (
	"fmt"
	"errors"
)

// CON errors.New()-------------------------------------------

func calcSalario (hrs float64, valorHr float64) (float64, error){
	if hrs < 80 {
		err := errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return 0, err
	}
	salario := hrs * valorHr
	if salario >= 150000 {
		salario = salario * 0.9
	}
	return salario, nil
}

func calcAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario <= 0 {
		err := errors.New("error: el salario no puede ser negativo")
		return 0, err
	}
	aguinaldo := mejorSalario / 12 * float64(mesesTrabajados)
	return aguinaldo, nil
}

func main() {
	salario , errS := calcSalario(60, 30)
	aguinaldo, errA := calcAguinaldo(salario, 6)

	if errS != nil {
		fmt.Println(errS)
	} else {
		fmt.Printf("El salario fue: %v\n", salario)
	}

	if errA != nil {
		fmt.Println(errA)
	} else {
		fmt.Printf("El aguinaldo fue: %v\n", aguinaldo)
	}

}