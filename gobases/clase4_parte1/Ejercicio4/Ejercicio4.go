package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horasTrabajadas int, valorHora float64) (salario float64, e error) {
	if horasTrabajadas < 80 {
		return -1, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salario = float64(horasTrabajadas) * valorHora
	if salario >= 150000 {
		salario -= (salario - 150000) * 0.1
	}
	return salario, nil
}

func calcularMedioAguinaldo(mejorSalario float64, mesesTrabajados int) (medioAguinaldo float64, e error) {
	if mejorSalario <= 0 || mesesTrabajados <= 0 {
		return -1, errors.New("error: el trabajador no puede haber tenido un ingreso menor o igual que 0 y/o trabajar menos que un mes")
	}
	medioAguinaldo = mejorSalario / 12 * float64(mesesTrabajados)
	return medioAguinaldo, nil
}

func main() {
	value, err := calcularMedioAguinaldo(80000, 6)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Println(value)
	}
	value2, err2 := calcularMedioAguinaldo(0, 6)
	if err2 != nil {
		println(err2.Error())
	} else {
		fmt.Println(value2)
	}
	value3, err3 := calcularMedioAguinaldo(80000, 0)
	if err3 != nil {
		println(err3.Error())
	} else {
		fmt.Println(value3)
	}

	value4, err4 := calcularSalario(80, 1000)
	if err4 != nil {
		println(err4.Error())
	} else {
		fmt.Println(value4)
	}
	value5, err5 := calcularSalario(10, 500)
	if err5 != nil {
		println(err5.Error())
	} else {
		fmt.Println(value5)
	}
}
