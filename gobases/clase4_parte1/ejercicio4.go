package main

import (
	"errors"
	"fmt"
)

func salarioMensual(horas int, valor float64) (float64, error) {
	if horas < 80 {
		return 0, errors.New("Error: El trabajador no puede haber trabajado menos de 80 horas mensuales")
	} else {
		var salario float64 = float64(horas) * valor
		if salario >= 150000 {
			salario = salario * 0.9
		}
		return salario, nil
	}
}

func medioAguinaldo(salario float64, meses int) (float64, error) {
	if salario < 0 || meses < 0 {
		return 0, errors.New("Error: No se pueden ingresar valores negativos")
	} else {
		var aguinaldo float64 = salario * float64(meses) / 12
		return aguinaldo, nil
	}
}

func main() {

	horas, valorHora, meses := 100, 10.5, 10
	salario, err := salarioMensual(horas, valorHora)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario para %d horas trabajadas es de %.2f\n", horas, salario)
	}

	aguinaldo, err := medioAguinaldo(salario, meses)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El aguinaldo para %d meses trabajados es de %.2f\n", meses, aguinaldo)
	}
}
