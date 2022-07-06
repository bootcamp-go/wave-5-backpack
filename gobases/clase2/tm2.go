// Ejercicio 2 - Calcular promedio

// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
package main

import (
	"errors"
	"fmt"
)

func promedio(notas ...int) (int, error) {
	var sumatoria int = 0
	for _, nota := range notas {
		if nota > 0 {
			sumatoria += nota
		} else {
			return 0, errors.New("hay un valor negativo\n")
		}
	}
	return sumatoria / len(notas), nil
}
func main() {
	prom, err := promedio(5, 6, 8, 5, 10, 9, 4)
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		fmt.Printf("El promedio es:%d\n", prom)

	}
}
