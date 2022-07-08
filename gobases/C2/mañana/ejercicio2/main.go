package main

import (
	"fmt"
)

/*Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. Se solicita generar una función en
la cual se le pueda pasar N cantidad de enteros y devuelva el promedio y un error en caso que uno de los números
ingresados sea negativo
*/

func calcularPromedio(notas ...float32) (float32, string) {
	var sumNotas float32
	for k, nota := range notas {
		if nota < 0 {
			return 0, fmt.Sprintf("Nota número %d con valor %.2f es inválida", k+1, nota)
		}
		sumNotas += nota
	}

	return sumNotas / float32(len(notas)), ""
}

func main() {
	promedio, msg := calcularPromedio(9, 9, -7, 9, 8)
	if msg != "" {
		fmt.Printf("%s\n", msg)
		return
	}

	fmt.Printf("El promedio de calificaciones es: %.2f\n", promedio)
}
