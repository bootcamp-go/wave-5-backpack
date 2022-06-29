package main

import (
	"errors"
	"fmt"
)

func calcPromedio(notas ...float32) (float32, error) {
	var totalSuma float32 = 0.0

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Error: Nota Negativa")
		}
		totalSuma += nota
	}
	return totalSuma / float32(len(notas)), nil
}

func main() {
	promedio, err := calcPromedio(5, 8, 6.5, 8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de notas del curso es: %.1f \n", promedio)
	}
}
