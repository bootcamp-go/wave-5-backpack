package main

import (
	"errors"
	"fmt"
)

type empleado struct {
	Nombre          string
	Apellido        string
	Salario         float64
	HsTrabajadas    int
	MesesTrabajados int
	Aguinaldo       float64
}

func (e empleado) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nSalario: 💰%.2f\nHs Trabajadas: %d\nMeses Trabajados: %d\nAguinaldo: 🤑%.2f", e.Nombre, e.Apellido, e.Salario, e.HsTrabajadas, e.MesesTrabajados, e.Aguinaldo)
}

func salarioTotal(HsTrabajadas int, valorHs float64) (float64, error) {
	salarioTotal := float64(HsTrabajadas) * valorHs

	if HsTrabajadas < 80 || HsTrabajadas < 0 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales. Hs ingresadas: %d", HsTrabajadas)
	} else if salarioTotal >= 150000 {
		salarioTotal -= (salarioTotal * 0.10)
		return salarioTotal, nil
	} else {
		return salarioTotal, nil
	}
}

func calcularAguinaldo(mesesTrabajados int, mejorSalario float64) (float64, error) {
	aguinaldo := (mejorSalario / 12) * float64(mesesTrabajados)

	if mesesTrabajados < 0 {
		return 0, errors.New("error: meses trabajados no puede ser un numero negativo")
	} else {
		return aguinaldo, nil
	}
}
