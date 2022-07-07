package main

import (
	"errors"
	"fmt"
)

func calcSalario(horasTrabajadas int, valorHora float64) (float64, error) {
	salario := float64(horasTrabajadas) * valorHora
	if salario >= 150000 {
		salario -= salario * 0.10
	}

	if horasTrabajadas < 80 || horasTrabajadas < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	return salario, nil
}

func main() {
	res, err := calcSalario(90, 5000)
	if err != nil {
		fmt.Println("Opps:", err)
	}
	fmt.Println(res)
}
