/*
	Ejercicio 4 - Calcular estadísticas

	Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de 
	calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, 
	máximo y promedio de sus calificaciones.

	Se solicita generar una función que indique qué tipo de cálculo se quiere realizar 
	(mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el 
	cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva 
	el cálculo que se indicó en la función anterior
*/

package main

import (
	"fmt"
	"errors"
)

const (
	min = "mínimo"
	prom = "promedio"
	max = "máximo"
)

func main() {
	operacion := calculoOperacion(min)
	resultado, err := operacion(7.0, 6.0, 3.0)

	if (err != nil ) {
		fmt.Println(err)
	} else {
		fmt.Printf("El cálculo de la operación es %.2v \n", resultado)
	}
}

func calculoOperacion(operacion string) func(valores ...float64) (float64, error) {
	var resultado float64
	switch operacion {
	case prom:
		return calculoPromedio(valores)
	case min:
		return calculoMinimo(valores)
	case max:
		return calculoMaximo(valores)
	}
	return resultado, nil
}

func calculoPromedio(valores []float64) (float64, error) {
	var prom float64
	for _, value := range valores {
		prom += value
	}
	if len(valores) == 0 {
		return 0, errors.New("Debe ingresar al menos dos números")
	}
	return prom/float64(len(valores)), nil
}

func calculoMinimo(valores []float64) (float64, error) {
	min := valores[0]
	for _, value := range valores {
		if value < min {
			min = value
		}
	}
	return min, nil
}

func calculoMaximo(valores []float64) (float64, error) {
	max := valores[0]
	for _, value := range valores {
		if value > max {
			max = value
		}
	}
	return max, nil
}

// PENDIENTE
