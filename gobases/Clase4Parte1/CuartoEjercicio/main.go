package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horas float64, valor float64) (float64, error) {
	horasTotal := horas * valor
	if horas > 150000 {
		horasTotal = horasTotal - horasTotal*10/100
	}
	if horas < 80 {
		err := errors.New("error: el trabajador no puede haber trabajado menos de 80hs mensuales")
		return 0, err
	}
	return horasTotal, nil
}

func calcularAguinaldo(sueldo float64, meses int) (float64, error) {
	if sueldo < 0 || meses < 0 {
		err := errors.New("error: el sueldo del trabajador o el numero de meses trabajados no pueden ser negativos")
		return 0, err
	} else {
		sueldoTotal := sueldo / 12 * float64(meses)
		return sueldoTotal, nil
	}

}

func main() {
	sueldoTrabajador, err := calcularSalario(90, 1000)
	if err != nil {
		fmt.Println(err)
	}
	aguinaldoTrabajador, err := calcularAguinaldo(sueldoTrabajador, 6)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sueldoTrabajador)
	fmt.Println(aguinaldoTrabajador)
}
