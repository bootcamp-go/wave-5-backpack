package main

import (
	"errors"
	"fmt"
)

type ErrorHoras struct{}

func (this *ErrorHoras) Error() string {
	return "error: el trabajador no puede trabajar menos de 80 hrs mensuales"
}

func salarioMensualTrabajador(horas int, valorHora float64) (float64, error) {
	salario := valorHora * float64(horas)

	if horas < 80 || horas < 0 {
		return 0, &ErrorHoras{}
	}

	if salario < 150000 {
		salario -= salario * 0.10
	}
	return salario, nil
}

func calcularAguinaldo(salariosSemestre []float64) (float64, error) {
	var mesesTrabajados int = len(salariosSemestre)

	if mesesTrabajados <= 0 {
		return 0, errors.New("error: el trabajador debe tener salarios para calcular el aguinaldo")
	}

	var mejorSalarioSemestre float64

	for i, salario := range salariosSemestre {
		if salario < 0 {
			err := fmt.Errorf("Salario ingresado en la posición %d: %f", i, salario)
			return 0, fmt.Errorf("error: el trabajador no puede tener salarios negativos. err: %w)", err)
		}

		if i == 0 {
			mejorSalarioSemestre = salario
		} else if salario > mejorSalarioSemestre {
			mejorSalarioSemestre = salario
		}
	}

	return mejorSalarioSemestre / 12 * float64(mesesTrabajados), nil
}

func main() {
	var valorHora float64 = 30000
	var horasTrabajadas []int = []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 8}
	var salariosSemestre []float64 = []float64{}

	for i, hora := range horasTrabajadas {
		salario, err := salarioMensualTrabajador(hora, valorHora)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			salariosSemestre = append(salariosSemestre, salario)
			fmt.Printf("El salario del trabajador en el mes %d es: %.1f$\n", i+1, salario)
		}
	}

	salariosSemestre = append(salariosSemestre, -1)

	aguinaldo, err := calcularAguinaldo(salariosSemestre)
	if err != nil {
		for err != nil {
			fmt.Println(err)
			err = errors.Unwrap(err)
		}
	} else {
		fmt.Println("El aguinaldo es:", aguinaldo)
	}

	fmt.Println()

	salariosSemestre = salariosSemestre[:len(salariosSemestre)-1]

	aguinaldo, err = calcularAguinaldo(salariosSemestre)
	if err != nil {
		for err != nil {
			fmt.Println(err)
			err = errors.Unwrap(err)
		}
	} else {
		fmt.Printf("El aguinaldo despues de quitar el causante del error es: %.1f$\n", aguinaldo)
	}

}
