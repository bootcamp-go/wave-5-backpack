package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horas int, valor_hora float64) (float64, error) {
	if horas < 80 {
		return 0, fmt.Errorf("Error: el trabajador no puede haber trabajado menos de 80h mensuales y trabajó %d.", horas)
	} else {
		salario := valor_hora * float64(horas)
		if salario >= 150000 {
			salario -= salario * 0.1
		}
		return salario, nil
	}
}

func medioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario < 0 || mesesTrabajados < 0 {
		return 0, errors.New("Error: no se pueden ingresar datos negativos.")
	}
	aguinaldo := (mejorSalario / 12) * float64(mesesTrabajados)
	return aguinaldo, nil
}

func main() {
	horas := 70
	valor_hora := 35000
	salario, err := calcularSalario(horas, float64(valor_hora))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario del trabajador es: %f", salario)
	}

	horas = 240
	valor_hora = 35000
	salario, err = calcularSalario(horas, float64(valor_hora))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario del trabajador es: %.2f\n", salario)
	}

}
