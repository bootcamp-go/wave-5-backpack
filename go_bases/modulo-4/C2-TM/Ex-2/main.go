package main

import (
	"errors"
	"fmt"
)

func main() {

	promedio, err := promAlumno(8, 7, 10, 9, 8, 6, -5)
	if err != nil {
		fmt.Println("No ingreses numeros negativos")
	} else {
		fmt.Printf("El promedio es: %v", promedio)
	}
}

func promAlumno(calis ...int) (float32, error) {
	var sumatoria int
	var promedio float32
	for _, v := range calis {
		if v < 0 {
			return 0, errors.New("No se aceptan numeros negativos")
		}
		sumatoria += v
	}
	fmt.Println(sumatoria)

	promedio = float32(sumatoria / len(calis))

	return promedio, nil
}
