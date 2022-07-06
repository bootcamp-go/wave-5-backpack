// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que se indicó en la función anterior

/* const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )

 ...

 minFunc, err := operation(minimum)
 averageFunc, err := operation(average)
 maxFunc, err := operation(maximum)

 ...

 minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
 averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
 maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
*/
package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(grades ...float64) float64 {
	var min float64

	for i, grade := range grades {
		if i == 0 || grade < min {
			min = grade
		}

	}
	return min
}
func averageFunc(grades ...float64) float64 {
	var sum float64 = 0

	for _, grade := range grades {
		sum += grade
	}

	return sum / float64(len(grades))
}
func maxFunc(grades ...float64) float64 {
	var max float64

	for i, grade := range grades {
		if i == 0 || grade > max {
			max = grade
		}

	}
	return max
}

func calculo(operador string) (func(...float64) float64, error) {
	switch operador {
	case minimum:
		return minFunc, nil
	case average:
		return averageFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		myError := fmt.Sprintf("el operador %s no existe", operador)
		return nil, errors.New(myError)
	}
}
func errorHandler(err error ) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
func main() {
	mini, err := calculo(minimum)
	errorHandler(err)
	maxi, err := calculo(maximum)
	errorHandler(err)
	averag, err := calculo(average)
	errorHandler(err)
	fmt.Printf("Minimo: %.2f \nMaximo: %.2f\nPromedio: %.2f", mini(6,8,9,7,10), maxi(6,8,9,7,10), averag(6,8,9,7,10))
}
