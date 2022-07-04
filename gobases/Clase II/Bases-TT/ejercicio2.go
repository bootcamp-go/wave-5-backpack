package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	alto    int
	ancho   int
	valores []float64
}

func (m Matrix) setData() {

	long := len(m.valores)
	nValores := m.alto * m.ancho

	if m.alto != m.ancho || nValores != long {
		fmt.Println("La matrix debe ser cuadrática, es decir, mismas filas y mismas columnas")
	} else {
		fmt.Println("Matrix ingresada correctamente")
	}
}

func (m Matrix) printMatrix() {

	for j := 0; j < m.ancho; j++ {
		fmt.Printf("%3.0f\n", m.valores[j*m.ancho:j*m.ancho+m.ancho])
	}
}

func (m Matrix) maxValue() {

	maximum := math.Inf(-1)

	for _, value := range m.valores {
		if value > maximum {
			maximum = value
		}
	}
	fmt.Println(maximum)
}

func main() {

	newMatrix := Matrix{alto: 3, ancho: 3, valores: []float64{10, 82, 33, 44, 102, 44, 85, 86, 150}}
	Matrix.setData(newMatrix)
	Matrix.printMatrix(newMatrix)
	Matrix.maxValue(newMatrix)

}
