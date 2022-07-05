package main

import (
	"fmt"
)

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
}

func (m Matrix) set() {
	if len(m.valores) != m.ancho*m.alto {
		fmt.Println("La cantidad de valores no coincide con las dimensiones de la matriz")
	}
}

func (m Matrix) esCuadratica() bool {
	var esCuadratica bool
	if (m.alto == m.ancho) && m.alto != 0 {
		esCuadratica = true
		return esCuadratica
	} else {
		esCuadratica = false
		return esCuadratica
	}

	return esCuadratica
}

func (m Matrix) print() {
	if len(m.valores) == 0 {
		fmt.Println("La matriz no tiene valores")
	}
	for fila := 0; fila < m.alto; fila++ {
		fmt.Printf("\t%.0f\n", m.valores[fila*m.ancho:fila*m.ancho+m.ancho])
	}
}

func main() {
	m := Matrix{
		valores: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		alto:    4,
		ancho:   4,
	}
	Matrix.set(m)
	Matrix.print(m)
	if Matrix.esCuadratica(m) {
		fmt.Println("Es cuadratica")
	} else {
		fmt.Println("No es cuadratica")
	}

}
