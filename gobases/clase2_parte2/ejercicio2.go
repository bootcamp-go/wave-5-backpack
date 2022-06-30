/*
	Ejercicio 2 - Matrix
	Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una 
	estructura que represente una matriz de datos.

	Para ello requieren una estructura Matrix que tenga los métodos:
	Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
	Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)

	La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión 
	del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main

import "fmt"

type Matrix struct {
	valores			[]float64
	alto			int
	ancho			int
}

func (m *Matrix) Set(valores ...float64) {
	m.valores = valores
	m.alto = 3
	m.ancho = 3
}

func (m Matrix) Print() {
	v := 0
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			fmt.Printf("%v",m.valores[v])
			v++
		}
		fmt.Println()
	}
}

func main() {
	m := Matrix{}
	m.Set(3,3,2,7,6,7,4,6,3)
	m.Print()
}