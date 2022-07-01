package main

import (
	"errors"
	"fmt"
)

func obtenerSalario(horasTrabajadas float64, valorDeHora float64) (float64, error) {

	if horasTrabajadas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	salario := horasTrabajadas * valorDeHora

	if salario > 150000 {
		salario *= 0.9
	}

	return salario, nil

}

func obtenerAguinaldo(salariosDelSemestre ...float64) (float64, error) {
	mesesTrabajados := len(salariosDelSemestre)

	if mesesTrabajados > 6 {
		return 0, fmt.Errorf("error: puede haber un máximo de 6 salarios en un semestre y se enviaron %d salarios", mesesTrabajados)
	}

	mejorSalario := .0

	for _, salario := range salariosDelSemestre {
		if salario < 0 {
			return 0, errors.New("error: no pueden existir salarios negativos")
		}
		if salario > mejorSalario {
			mejorSalario = salario
		}
	}

	return mejorSalario / 12 * float64(mesesTrabajados), nil
}

func main() {

	salario, err := obtenerSalario(78, 156.50)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Salario: %.2f\n", salario)
	}

	salario2, err := obtenerSalario(82, 156.50)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Salario2: %.2f\n", salario2)
	}

	aguinaldo, err := obtenerAguinaldo(196000, 197000, 198000, 199000, 199999, 200000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Aguinaldo: %.2f\n", aguinaldo)
	}

	aguinaldo2, err := obtenerAguinaldo(196000, 197000, 198000, 199000, 199999, 200000, 201000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Aguinaldo2: %.2f\n", aguinaldo2)
	}

	aguinaldo3, err := obtenerAguinaldo(196000, -197000, 198000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Aguinaldo2: %.2f\n", aguinaldo3)
	}
}
