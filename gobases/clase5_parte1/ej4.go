package main

import (
	"errors"
	"fmt"
)

const SALARIO_MINIMO = 150000

func main() {
	salario, err := salarioMensual(100, 4000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Salario: ", salario)

	medioAguinaldo, err := medioAguinaldo(500000, 6)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Medio aguinaldo: ", medioAguinaldo)
}

func salarioMensual(horasTrabajadas int, precioPorHora float64) (float64, error) {

	if horasTrabajadas < 80 {
		return 0, errors.New("El trabajador no pudo haber trabajado menos de 80 hs mensuales")
	}

	salario := float64(horasTrabajadas) * precioPorHora

	if salario >= SALARIO_MINIMO {
		return salario * 0.9, nil
	}

	return salario, nil
}

func medioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario <= 0 || mesesTrabajados <= 0 {
		return 0, errors.New("Datos inválidos ingresados")
	}

	resultado := (mejorSalario * float64(mesesTrabajados)) / 12

	return resultado, nil
}
