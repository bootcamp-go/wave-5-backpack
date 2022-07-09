package main

import (
	"fmt"
	"math"
)

//Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
//Para ello requieren una estructura Matrix que tenga los métodos:
//Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
//Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
//La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

type Matrix struct {
	values []float64
	height int
	width  int
}

func (m Matrix) set() {
	if len(m.values) != m.height*m.width {
		fmt.Println("El número de valores ingresados no concuerda con los parámetros solicitsdos")
	}
}

func (m Matrix) Quadratic() bool {
	if m.height == m.width && m.height != 0 {
		return true
	}
	return false
}

func (m Matrix) Print() {
	if len(m.values) == 0 {
		fmt.Println("La matriz no contiene valores")
	}
	for i := 0; i < m.height; i++ {
		fmt.Printf("\t%.0f\n", m.values[i*m.width:i*m.width+m.width])
	}
}

func (m Matrix) Max() float64 {
	max := -math.MaxFloat64
	for _, value := range m.values {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	m := Matrix{values: []float64{5, 6, 3, 8, 9, 2}, height: 3, width: 3}
	Matrix.set(m)
	Matrix.Print(m)
	Matrix.Quadratic(m)
	Matrix.Max(m)

}
