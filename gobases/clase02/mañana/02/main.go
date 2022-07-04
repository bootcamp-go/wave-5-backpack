// Ejercicio 2 - Calcular promedio
// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

package main
import (
	"fmt"
	"errors"
)

func calcAverage (valores ... float64) (float64, error) {
	suma := 0.0
	for _, value := range valores {
		if value < 0 {
			return 0, errors.New("No puede haber valores negativos")
		}
		suma += value
	}
	return suma/float64(len(valores)), nil
}

func main(){
	fmt.Println(calcAverage(7,-8,3,4.5))
}