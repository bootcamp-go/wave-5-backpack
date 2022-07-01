package main

/* Ejercicio 2 - Matrix

Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
	1. Set: Recibe una serie de valores de punto flotante e inicializa los valores en la
estructura Matrix
	2. Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea
entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la
dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

import (
	"fmt"
	"math"
)

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func (m Matrix) Set() {
	if len(m.valores) != m.ancho*m.alto {
		fmt.Println("La cantidad de valores no coincide con las dimensiones especificadas")
	}

}
func (m Matrix) Cuadratica() bool {
	if (m.alto == m.ancho) && m.alto != 0 {
		return true
	}
	return false
}
func (m Matrix) Max() float64 {
	max := -math.MaxFloat64
	for _, elemento := range m.valores {
		if elemento > max {
			max = elemento
		}
	}
	return max
}
func (m Matrix) Print() {
	if len(m.valores) == 0 {
		fmt.Println("La matriz está vacía")
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Printf("\t%.0f\n", m.valores[fila*m.ancho:fila*m.ancho+m.ancho])
	}
}
func main() {
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		alto:    3,
		ancho:   3,
	}
	Matrix.Set(m)
	Matrix.Print(m)
	fmt.Println(Matrix.Cuadratica(m))
}
