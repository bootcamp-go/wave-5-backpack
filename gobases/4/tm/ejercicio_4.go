package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	horas := 160
	valor := 1000
	meses := -6

	salario, err := SalarioMensual(horas, valor)
	if err != nil {
		err1 := fmt.Errorf("[ERROR]: %w", err)
		fmt.Println(err1)
		fmt.Println(errors.Unwrap(err1))
		os.Exit(1)
	}

	aguinaldo, err := MedioAguinaldo(salario, meses)
	if err != nil {
		err1 := fmt.Errorf("[ERROR]: %w", err)
		fmt.Println(err1)
		fmt.Println(errors.Unwrap(err1))
		os.Exit(1)
	}

	fmt.Printf("salario mensual: %.2f\naguinaldo: %2.f\n", salario, aguinaldo)
}

func SalarioMensual(horas int, valor int) (salario float64, err error) {
	salario = float64(horas) * float64(valor)
	if horas <= 80 || horas < 0 {
		return 0, errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}

	if salario >= 150000 {
		return salario * 0.9, nil
	} else {
		return salario, nil
	}
}

func MedioAguinaldo(mayor_salario float64, meses int) (aguinaldo float64, err error) {
	if meses < 0 {
		return 0, errors.New("no se admiten numeros negativos")
	}

	aguinaldo = mayor_salario / 12 * float64(meses)
	return aguinaldo, nil
}
