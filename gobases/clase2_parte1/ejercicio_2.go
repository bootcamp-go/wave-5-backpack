package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(calificaciones ...int) (float32, error) {
	suma := 0
	cantidad := len(calificaciones)

	for _, valor := range calificaciones {
		if valor < 0 {
			return 0, errors.New("La calificación no puede ser negativa.")
		} else {
			suma = suma + valor
		}
	}

	promedio := float32(suma) / float32(cantidad)

	return promedio, nil
}

func main() {
	res, err := calcularPromedio(2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio del estudiante es ", res)
	}

	res2, err2 := calcularPromedio(2, -3, 4, 5)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("El promedio del estudiante es ", res2)
	}
}
