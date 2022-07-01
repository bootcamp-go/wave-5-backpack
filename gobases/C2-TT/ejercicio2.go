package main

import "fmt"

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	cuadratica bool
	valorMax   int
}

func (m *Matrix) Set(valores ...float64) {
	m.valores = valores
}

func (m Matrix) Print() {
	for i := 0; i < m.alto; i++ {
		for j := 0; j <= m.ancho; j++ {
			fmt.Printf("%v\t", m.valores[i+(j*m.ancho)])
		}
		fmt.Printf("\n")
	}
}

func main() {
	matrix := Matrix{
		alto:  3,
		ancho: 2,
	}

	matrix.Set(1, 2, 3, 4, 5, 6, 7)

	matrix.Print()
}
