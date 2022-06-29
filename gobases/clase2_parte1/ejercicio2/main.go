package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float64) (float64, error) {
	var total float64
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("La nota no puede ser negativa")
		}

		total += nota
	}

	return total / float64(len(notas)), nil
}

func main() {
	fmt.Println(calcularPromedio(5.0, 5.0, 5.0, 5.0))
	fmt.Println(calcularPromedio(5.0, 4.2, 3.7, 1.0, 3.6, 4.7))
}
