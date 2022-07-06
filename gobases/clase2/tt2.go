// Ejercicio 2 - Matrix
// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	height int
	width  int
	values []float64
}

func (mat Matrix) Set() {
	if len(mat.values) != mat.height * mat.width {
		fmt.Printf("Faltan valores para el tamaño de la matriz")
	}

}
func (mat Matrix) Cuadratica() bool {
	if (mat.height == mat.width) && mat.height != 0 {
		return true
	}
	return false
}
func (mat Matrix) MaxValue() float64 {
	max := -math.MaxFloat64
	for _, num := range mat.values {
		if num > max {
			max = num
		}
	}
	return max
}

func (mat Matrix) Printer() {
	if len(mat.values) == 0 {
		fmt.Println("La matriz está vacía")
	} else {
		for fila := 0; fila < mat.height; fila++ {
			fmt.Printf("\t%.0f\n", mat.values[fila*mat.width:fila*mat.width+mat.width])
			//fmt.Println("aca")
		}
	}
}
func main() {
	matriz := Matrix{
		values: []float64{2, 3, 4, 5, 6, 7, 8, 9,1},
		height: 3,
		width:  3,
	}
	Matrix.Set(matriz)
	Matrix.Printer(matriz)
	fmt.Printf("Es cuadrada?: %v\nEl valor maximo es: %.2f\n", Matrix.Cuadratica(matriz), Matrix.MaxValue(matriz))
}
