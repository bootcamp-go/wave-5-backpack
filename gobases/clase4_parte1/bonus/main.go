package main

import (
	"errors"
	"fmt"
)

const (
	MIN_HORAS          = 80
	MIN_SALARIO_IMP    = 150000
	PORC_DESCUENTO_IMP = 10
)

func descontarImpuestos(salario float64) (float64, error) {
	if salario <= 0 {
		return 0, errors.New("El salario ingresado debe ser mayor a cero")
	}
	if salario >= MIN_SALARIO_IMP {
		return (salario - (salario * PORC_DESCUENTO_IMP / 100)), nil
	}
	return salario, nil
}

func calcularSalario(horas int, valorHora float64) (float64, error) {
	if horas < MIN_HORAS {
		return 0, fmt.Errorf("el trabajador no puede haber trabajado menos de %d hrs mensuales", MIN_HORAS)
	}
	salarioTotal, err := descontarImpuestos(valorHora * float64(horas))
	if err != nil {
		return 0, fmt.Errorf("Error en el sistema : %w", err)
	}
	return salarioTotal, nil
}

func calcularSalarioAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {
	if mejorSalario <= 0 || mesesTrabajados <= 0 {
		return 0, errors.New("Esta funcion no acepta valores negativos")
	}
	return (mejorSalario / 12 * float64(mesesTrabajados)), nil
}

func main() {
	//Salario
	salario, err := calcularSalario(90, 1)

	if err != nil {
		err1 := errors.Unwrap(err)
		if err1 != nil {
			fmt.Println("Errores anidados")
		}
		fmt.Println(err)
		return
	}
	fmt.Println("El salario es de :", salario)

	//Aguinaldo
	aguinaldo, err := calcularSalarioAguinaldo(100, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El aguinaldo es de :", aguinaldo)
}
