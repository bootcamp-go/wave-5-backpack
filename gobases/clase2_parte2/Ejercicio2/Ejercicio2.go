package main

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
		fmt.Println("La matriz debe llenarse y por ello debe cumplir que la cantidad de valores es igual al ancho por el largo de la matriz")
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
		fmt.Println("Matriz vacia")
	}
	aux := 0
	for i := 0; i < m.ancho; i++ {
		for j := 0; j < m.alto; j++ {
			fmt.Printf("%v\t", m.valores[aux])
			aux++
		}
		fmt.Println()
	}
}
func main() {
	m1 := Matrix{
		valores: []float64{3, 7, 9, 1, 6, 0, 2, 5, 3, 7, 9, 1, 6, 0, 2, 5},
		alto:    4,
		ancho:   4,
	}

	m2 := Matrix{
		valores: []float64{16, 8, 5, 9, 2, 34, 5, 0, 12},
		alto:    3,
		ancho:   3,
	}
	Matrix.Set(m1)
	fmt.Println(Matrix.Cuadratica(m1))
	Matrix.Print(m1)
	Matrix.Set(m2)
	fmt.Println(Matrix.Cuadratica(m2))
	Matrix.Print(m2)
}
