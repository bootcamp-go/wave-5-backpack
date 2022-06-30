/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

*/

package main

import (
	"fmt"
)

func minimo(valores...float64) float64{
	min := valores[0]
	for _, valor := range valores {
		if valor < min {
			min = valor
		}
	}
	return min
}

func maximo(valores...float64) float64{
	max := valores[0]
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	return max
}

func promedio (valores ... float64) float64{
	suma := 0.0
	for _, value := range valores {
		/* if value < 0 {
			return 0, errors.New("No puede haber valores negativos")
		} */
		suma += value
	}
	return suma/float64(len(valores))
}

func calculo (tipoDeCalculo string, valores ... float64) float64 {
	switch tipoDeCalculo{
	case "minimo":
		return minimo(valores...)
	
	case "maximo":
		return maximo(valores...)
	
	case "promedio":
		return promedio(valores...)
	}
	return 0
}

func main(){
	fmt.Println(calculo("ssss", 3,4,5,6))
}