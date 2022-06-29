package main

import (
	"errors"
	"fmt"
)

// Ejercicio 2 - Calcular promedio

// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.

// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error
// en caso que uno de los números ingresados sea negativo

func calcularPromedio(calificaciones ...float64) (float64, error) {
	sumatoria := 0.0
	totalMaterias := 0

	for _, v := range calificaciones {
		if v < 0 {
			return 0, errors.New("No puede haber calificaciones negativas")
		}
		sumatoria += v
		totalMaterias++
	}

	return sumatoria / float64(totalMaterias), nil
}

func main() {
	fmt.Println("Ejercicio 2 - Calcular promedio")
	fmt.Println("")

	promedioAlumno1, err := calcularPromedio(8, 5, 6, 10, 9, 7)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Promedio del alumno 1: ", promedioAlumno1)
	}

	promedioAlumno2, err := calcularPromedio(8, 5, -6, 10, 9, 7)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Promedio del alumno 2: ", promedioAlumno2)
	}
}
