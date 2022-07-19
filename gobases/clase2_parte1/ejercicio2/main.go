package main

import (
	"errors"
	"fmt"
)

func main() {
	calificaciones := []float32{4.1, 4.1, 4.1}
	fmt.Println(calificaciones)
	promedio, err := promediarCalificaciones(calificaciones...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio de las notas es: ", promedio)
	}
}
func promediarCalificaciones(calificaciones ...float32) (float32, error) {
	var totalNotas float32
	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0, errors.New("Las notas no deben ser negativas")
		}
		totalNotas += float32(calificacion)
	}
	return totalNotas / float32(len(calificaciones)), nil
}
