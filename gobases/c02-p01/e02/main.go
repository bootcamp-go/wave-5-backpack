package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...int) (float64, error) {
	suma := 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Ingresaste un valor incorrecto")
		}
		suma += nota
	}
	return float64(suma) / float64(len(notas)), nil
}

func main() {

	fmt.Println("Las notas son:", 10, 9, 8, 7, 6)
	prom01, err01 := calcularPromedio(10, 9, 8, 7, 6)

	if err01 != nil {
		fmt.Println(err01)
	} else {
		fmt.Println("El promedio es: ", prom01)
	}

	fmt.Println("Las notas son:", 10, 9, 8, 7, -6)
	prom02, err02 := calcularPromedio(10, 9, 8, 7, -6)

	if err02 != nil {
		fmt.Println(err02)
	} else {
		fmt.Println(prom02)
	}

}
