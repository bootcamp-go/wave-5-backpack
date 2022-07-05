package main

import (
	"errors"
	"fmt"
)

//Bonus Track -  Impuestos de salario #4

func salMens(horas int, valor float64) (float64, error) {
	sal := valor * float64(horas)
	if horas < 80 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if sal >= 150000 {
		desc := 0.10 * sal
		sal -= desc
		return sal, nil
	}
	return sal, nil
}

func medAgui(bestSal float64, mesTrab int) (float64, error) {
	aguinaldo := (bestSal / 12) * float64(mesTrab)
	if bestSal < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	return aguinaldo, nil
}

func evaluador(horas int, valor float64, mesesTrabajado int) {
	val, err := salMens(horas, valor)
	fmt.Println("Su salario es", val)

	//aguinaldo prueba
	agui, err1 := medAgui(val, mesesTrabajado)
	err1 = fmt.Errorf("Error salario: %w", err)
	if errors.Unwrap(err1) != nil {
		fmt.Println(errors.Unwrap(err1))
	}
	fmt.Println("Su aguinaldo es", agui)
}

func main() {
	//Indicar horas Trabajadas, Valor de la hora, Meses trabajados en el semestre
	evaluador(78, 4000, 5)
}
