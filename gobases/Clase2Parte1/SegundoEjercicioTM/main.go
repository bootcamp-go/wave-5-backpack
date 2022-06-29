package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float64) (float64, error) {
	var resultado float64
	for _, value := range notas {
		if value > 0 {
			resultado += value
		} else {
			return 0, errors.New("Las notas no pueden ser negativas\n")
		}
	}
	return resultado / float64(len(notas)), nil
}

func main() {
	promedioNotas, error := calcularPromedio(5, 2, 3.5, 10, 5)

	if error != nil {
		fmt.Printf("Hubo un error en el calculo\n")
	} else {
		fmt.Printf("Promedio del estudiante: %v \n", promedioNotas)
	}
}
