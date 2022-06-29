package main

import "errors"

func calcularPromedio(calificaciones ...int) (float64, error) {
	//Ejercicio 2
	contNotas := 0
	sumaNotas := 0
	for _, calificacion := range calificaciones {

		if calificacion < 0 {
			return 0, errors.New("La calificacion no puede ser negativa")
		}
		sumaNotas += calificacion
		contNotas++
	}

	promedio := sumaNotas / contNotas

	return float64(promedio), nil
}
