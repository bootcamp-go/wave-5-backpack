/*
Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores 

mínimo, 
máximo y 
promedio de sus calificaciones.

Se solicita generar una función que indique: 
a) qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y 
b) que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

*/

package main


import (
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )
 

func main() {

	minFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

}

func operation(operador string) func(valor1, valor2 float64) float64 {
	switch operador {
	case minimum:
		return opMinimum
	case maximum:
		return opMaximum
	case average:
		return opMaximum
	}
	return nil
}

func opMinimum(calificaciones ... float64) float64 {
// For para buscar el minimo


	return valor1 + valor2
}

func opMaximum(valor1, valor2 float64) float64 {
	return valor1 - valor2
}

func opAverage(valor1, valor2 float64) float64 {
	return valor1 * valor2
}


func calcularEstadisticas (minutosTrabajados int, categoria string) float64 {

	var horasTrabajadas float64 = float64(minutosTrabajados) / 60
	var salario float64 = 0.0
	switch categoria {
    case "A":
		//Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual
        salario = (3.000 * horasTrabajadas) * 1.5
    case "B":
		// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
		salario = (1.500 * horasTrabajadas) * 1.2
    case "C":
        salario = (1.000 * horasTrabajadas)   
    }

	return salario
}

