package main

import (
	"errors"
	"fmt"
)

//Ejercicio 2 - calcular promedio

func promAlumno(valores ...int) (float64, error) {
	var cont int = 0
	var ke int = 0
	for i, value := range valores {
		if value < 0 {
			return 0, errors.New("Una de las notas es negativa")
		}
		cont += value
		ke = i + 1
	}
	return float64(cont) / float64(ke), nil
}

func main() {
	prom, er := promAlumno(4, -5, 3, 2, 1)
	fmt.Println("Su promedio es: ", prom)
	fmt.Println("Su Error es: ", er)
}
