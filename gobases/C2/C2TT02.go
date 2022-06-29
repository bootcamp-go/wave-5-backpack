package main

import (
	"fmt"
)

type Matrix struct {
	valores []float64
	alto    int
	ancho   int
	maximo  float64
}

func (m *Matrix) Set(valores ...float64) {
	m.valores = valores
}

func (m Matrix) Print() {

	for j := 0; j < m.alto; j++ {
		for i := 0; i < m.ancho; i++ {
			fmt.Printf("%v\t", m.valores[i+(j*2+j)])
			//fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func main() {
	matrix := Matrix{
		alto:  2,
		ancho: 4,
	}

	matrix.Set(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)

	if matrix.alto == matrix.ancho {
		fmt.Println("La matriz es cuadratica")
	} else {
		fmt.Println("La matriz no es cuadratica")
	}

	matrix.Print()
}
