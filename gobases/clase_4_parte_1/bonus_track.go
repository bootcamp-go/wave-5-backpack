package main

import (
	"errors"
	"fmt"
)

const MINIMO_HORAS = 80
const MEJOR_SALARIO = 500000
const MESES_DE_TRABAJO = 5

func calcularPorcentage(numero int, porcentage int) int {
	return int(float32(numero) * (float32(porcentage) / float32(100)))
}

func calcularSalario(horas int, precio int) (int, error) {

	if horas < MINIMO_HORAS {

		return 0, errors.New("error: el trabajador no puede haber trabajado menos de " + fmt.Sprint(MINIMO_HORAS) + " hs mensuales")
	}

	salario := horas * precio
	if salario > 150000 {
		salario += calcularPorcentage(salario, 10)
	}
	return horas * precio, nil
}

func calcularAguinaldo(mejorSalarioSemestre int, mesesTrabajados int) (int, error) {
	var err error
	if mejorSalarioSemestre <= 0 {
		err = fmt.Errorf("error: se ingreso como mejor salario un número negativo: %d", mejorSalarioSemestre)
	}

	if mesesTrabajados <= 0 {
		err = fmt.Errorf("error: se ingreso en meses de trabajo un número negativo o 0: %d \n %w", mejorSalarioSemestre, err)
	}

	if err != nil {
		return 0, err
	}
	return mejorSalarioSemestre / (12 * mesesTrabajados), nil
}

func main() {
	salario, errSalario := calcularSalario(2, 5000)

	if errSalario != nil {
		fmt.Println(errSalario)
	}
	println("Salario:", salario)

	aguinaldo, errAguinaldo := calcularAguinaldo(500000, 3)
	if errAguinaldo != nil {
		if errors.Unwrap(errAguinaldo) != nil {
			fmt.Println(errors.Unwrap(errAguinaldo))
		}
		fmt.Println(errAguinaldo)
	}
	fmt.Println("Aguinaldo: ", aguinaldo)
}
