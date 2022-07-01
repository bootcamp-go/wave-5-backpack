package main

import "fmt"

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

type Matrix struct {
	valores []float32
	alto    int
	ancho   int
}

func (m Matrix) Set() {
	if len(m.valores) != m.alto*m.ancho {
		fmt.Println("los valores no coinciden")
		panic("No se puede continuar con el programa")
	}
}

func (m Matrix) Print() {

	if len(m.valores) == 0 {
		fmt.Println("La matriz está vacía")
	}
	for i := 0; i < m.alto; i++ {
		fmt.Printf("\t%.0f\n", m.valores[i*m.ancho:i*m.ancho+m.ancho])
	}
}

func (m Matrix) Cuadratica() bool {
	if (m.alto == m.ancho) && m.alto != 0 {
		return true
	}
	return false
}

func (m Matrix) Max() float32 {
	var maximo float32

	for _, matriz := range m.valores {
		if maximo < matriz {
			maximo = matriz
		}

	}

	return maximo

}

func main() {
	matriz := Matrix{
		valores: []float32{1, 2, 3, 4, 5, 6, 7, 8},
		alto:    3,
		ancho:   3,
	}
	Matrix.Set(matriz)
	Matrix.Print(matriz)
	fmt.Println(Matrix.Cuadratica(matriz))
}
