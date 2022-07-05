package main

import (
	"errors"
	"fmt"
)

func CalcularSalario(horas_trabajadas int, precio_hora float64) (float64, error) {
	if horas_trabajadas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salario := precio_hora * float64(horas_trabajadas)
	if salario > 150000 {
		return salario * (1 - 0.1), nil
	}
	return salario, nil
}

func CalcularAguinaldo(mejor_salario_semestre float64, meses_trabajados int) (float64, error) {
	if mejor_salario_semestre < 0 || meses_trabajados < 0 {
		return 0, fmt.Errorf("no pueden ingresarse valores negativos")
	}
	return mejor_salario_semestre / 12 * float64(meses_trabajados), nil
}

func main() {
	calculo, err := CalcularSalario(80, 10000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(calculo)
	aguinaldo, err2 := CalcularAguinaldo(150000, 12)
	if err2 != nil {
		fmt.Println(errors.Unwrap(err2))
		return
	}
	fmt.Println(aguinaldo)
}
