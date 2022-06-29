package main

import (
	"errors"
	"fmt"
)

func promedioDeCalificaciones(calificaciones ...int) (promedio float64, err error) {
	suma := 0
	err = nil
	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			err = errors.New("No se permite ingresar calificaciones negativas")
			break
		}
		suma += calificacion
	}
	promedio = float64(suma) / float64(len(calificaciones))
	return
}

func main() {
	promedio, err := promedioDeCalificaciones(5, 7, 3, 8, 10)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Promedio: %f\n", promedio)
	}
}
