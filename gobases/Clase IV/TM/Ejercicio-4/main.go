package main

import (
	"errors"
	"fmt"
)

// Estructura de error

type errorInitial struct{}

//Método de estructura error

func (e errorInitial) Error() string {
	return "Ha ocurrido un error"
}

// Función para determinar salario con hh trabajadas y valor de hh.

func calcularSalario(hh float64, vhh float64) (float64, error) {

	salario := hh * vhh

	if hh < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80hrs mensuales")
	} else if salario >= 150000 {
		salario = salario * 0.9
	}
	return salario, nil
}

// Función para determinar aguinaldo

func calcularMedioAguinaldo(salariosSemestre ...float64) (float64, error) {

	mesesTrabajados := len(salariosSemestre)
	var mejorSalario float64 = 0.0

	for _, salario := range salariosSemestre {

		if salario < 0 {
			return 0, fmt.Errorf("error, se ha ingresado un salario negativo: %0.1f", salario)
		} else if salario > mejorSalario {
			mejorSalario = salario
		}
	}

	aguinaldo := mejorSalario / (12.0 * float64(mesesTrabajados))

	return aguinaldo, nil

}

func main() {

	salario, err1 := calcularSalario(80, 30000)
	if err1 != nil {
		fmt.Println(errors.Unwrap(err1))
	} else {
		fmt.Printf("El salario para este mes es: %.1f \n", salario)
	}

	aguinaldo, err2 := calcularMedioAguinaldo(1000.0, 3000.0, 4450.0, 5500.0, 6500.0, 9600.0)
	if err2 != nil {
		fmt.Println(errors.Unwrap(err2))
	} else {
		fmt.Printf("El aguinaldo medio es: %.2f \n", aguinaldo)
	}

}
