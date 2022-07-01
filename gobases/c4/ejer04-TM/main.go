package main

import (
	"errors"
	"fmt"
)

// Vamos a hacer que nuestro programa sea un poco más complejo.
// Desarrolla las funciones necesarias para permitir a la empresa calcular:
// Salario mensual de un trabajador según la cantidad de horas trabajadas.
// La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
// Dicha función deberá retornar más de un valor (salario calculado y error).
// En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
// En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
// El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
// Calcular el medio aguinaldo correspondiente al trabajador
// Fórmula de cálculo de aguinaldo:
// [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
// La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

const (
	MIN_HORAS     = 80
	SALARIOMINIMP = 150000
	IMP           = 10
)

type errorPersonalizado struct {
	mensaje string
}

type errorUno struct{}
type errorDos struct{}

func (e errorUno) Error() string {
	return "error 1"
}

func (e errorDos) Error() string {
	return "error 2"
}

func (c *errorPersonalizado) Error() string {
	return fmt.Sprintf(c.mensaje)
}

func calcularSalario(horas int, valorHora float64) (float64, error) {
	if horas < MIN_HORAS {
		return 0.0, &errorPersonalizado{"el trabajador no puede haber trabajado menos de 80 hrs mensuales"}
	}
	salarioTotal, err := descontarImpuestos(valorHora * float64(horas))
	if err != nil {
		return 0, fmt.Errorf("Error en el sistema : %w", err)
	}
	return salarioTotal, nil
}

func descontarImpuestos(salario float64) (float64, error) {
	if salario <= 0 {
		return 0, errors.New("El salario ingresado debe ser mayor a cero")
	}
	if salario >= SALARIOMINIMP {
		return (salario - (salario * IMP / 100)), nil
	}
	return salario, nil
}

func calcularAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario <= 0 || mesesTrabajados <= 0 {
		return 0, errors.New("No ingrese valores negativos")
	}
	return (mejorSalario / 12 * float64(mesesTrabajados)), nil
}

func main() {

	e1 := errorUno{}
	e2 := fmt.Errorf("errores: %w", e1)
	e3 := fmt.Errorf("error mas %w", e2)

	salario, err := calcularSalario(60, 1)

	if err != nil {
		fmt.Println(err)
		errors.Unwrap(e3)
		err3 := fmt.Errorf("%w", e3)
		fmt.Println(err3)
		return
	}
	fmt.Println("El salario es:", salario)

	aguinaldo, err := calcularAguinaldo(200000, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El aguinaldo es:", aguinaldo)
}
