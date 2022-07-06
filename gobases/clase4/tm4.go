/*Ejercicio 4 -  Impuestos de salario #4



Vamos a hacer que nuestro programa sea un poco más complejo y útil.
	1. Desarrolla las funciones necesarias para permitir a la empresa calcular:
		a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
			1. La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
			2. Dicha función deberá retornar más de un valor (salario calculado y error).
			3. En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar
			el 10% en concepto de impuesto.
			4. En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
			la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber
			trabajado menos de 80 hs mensuales”.
		b) Calcular el medio aguinaldo correspondiente al trabajador
			1. Fórmula de cálculo de aguinaldo:
			[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
			2. La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en
			caso de que se ingrese un número negativo.

	2. Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
	“fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu
	función “main()”.
	REVEER PORQUE ES SOLUCION DE CLASE LA MIA NO SALIO
	
	*/


package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e MyError) Error() string {
	return "Los valores ingresados son incorrectos"
}

func Salary(horasTrabajadas int, precioHora float64) (float64, error) {
	var salary float64
	salary = float64(horasTrabajadas) * precioHora
	if horasTrabajadas < 80 || horasTrabajadas < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	if salary >= 150000 {
		impuesto := (salary * 10 / 100)
		salary -= impuesto
		return salary, nil
	}

	return salary, nil
}

func CalcularMedioAguinaldo(salary float64, mesesTrabajados int) (float64, error) {
	//[mejor salario del semestre] / 12 * [meses trabajados en el semestre]
	if mesesTrabajados < 0 {
		return 0, fmt.Errorf("El trabajador no tiene meses trabajados. Valor %d inválido", mesesTrabajados)
	}

	aguinaldo := salary / float64(12) * float64(mesesTrabajados)
	return aguinaldo, nil
}

func main() {
	var precioHora float64 = 2500
	var horasTrabajadas int = 80
	var mesesTrabajados int = 6

	sueldo, err := Salary(horasTrabajadas, precioHora)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Sueldo 💰", sueldo)

	mError := &MyError{}
	e1 := fmt.Errorf("mError: %w", mError)
	aguinaldo, err := CalcularMedioAguinaldo(sueldo, mesesTrabajados)
	if err != nil {
		fmt.Println(errors.Unwrap(e1))
		return
	}

	fmt.Println("Aguinaldo 🤑 $", aguinaldo)
}
