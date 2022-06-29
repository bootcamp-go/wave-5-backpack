package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calcularPromedio(8.0, 3.5, 9.4, 6.6)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de las notas es de %.2f\n", promedio)
	}
}

func calcularPromedio(notas ...float64) (float64, error) {
	var promedio float64
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Una nota no puede tener un valor negativo")
		}
		promedio += nota
	}
	return promedio / float64(len(notas)), nil
}
