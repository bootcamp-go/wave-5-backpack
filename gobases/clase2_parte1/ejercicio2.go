/*
	Ejercicio 2 - Calcular promedio

	Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. 
	Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros 
	y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

package main

import (
	"fmt"
	"errors"
)

func main() {
	resultado, err := calculoPromedio(5, 7, 7, 6, 3, 1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio es: %.2v \n", resultado)
	}
}

func calculoPromedio(calificaciones ...float64) (float64, error) {
	var promedio float64
	for _, value := range calificaciones {
		if value < 0 {
			return 0, errors.New("Uno de los números ingresados es negativo.")
		} 
		promedio += value
	}
	if len(calificaciones) == 0 {
		return promedio, nil
	}
	return promedio/float64(len(calificaciones)), nil
}