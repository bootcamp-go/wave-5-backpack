package main

import (
	"errors"
	"fmt"
)

// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio
// y un error en caso que uno de los números ingresados sea negativo

func main() {
	promedio, err := calcularPromedio(4, 7, 9, 7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Promedio alumno 1: %v\n", promedio)
	}

	promedio2, err := calcularPromedio(5, 8, 4, 9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Promedio alumno 2: %v\n", promedio2)
	}

	promedio3, err := calcularPromedio(-1, 5, 8, 10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Promedio alumno 3: %v\n", promedio3)
	}
}

func calcularPromedio(notas ...int) (float64, error) {
	total := len(notas)
	sum := 0

	for _, v := range notas {
		if v < 0 {
			return 0, errors.New("No puede ingresar número negativo")
		}

		sum += v
	}

	promedio := float64(sum) / float64(total)

	return promedio, nil
}
