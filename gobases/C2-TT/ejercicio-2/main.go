package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	valores      []float64
	alto         int
	ancho        int
	cuadratica   bool
	valor_maximo float64
}

func (matrix *Matrix) Set(valores ...float64) {
	if matrix.alto == matrix.ancho {
		matrix.cuadratica = true
	} else {
		matrix.cuadratica = false
	}

	matrix.valores = make([]float64, matrix.alto*matrix.ancho)

	for i := 0; i < matrix.alto*matrix.ancho; i++ {
		if i < len(valores) {
			matrix.valores[i] = valores[i]
		} else {
			matrix.valores[i] = 0.0
		}
	}

	max := -math.MaxFloat64
	for _, elemento := range matrix.valores {
		if elemento > max {
			max = elemento
		}
	}
	matrix.valor_maximo = max

}

func (matrix Matrix) Print() {
	for fila := 0; fila < matrix.alto; fila++ {
		fmt.Printf("\t%.0f\n", matrix.valores[fila*matrix.ancho:fila*matrix.ancho+matrix.ancho])
	}
	fmt.Printf("Valor maximo: %.2f\n", matrix.valor_maximo)
	fmt.Printf("Cudratica: %t\n", matrix.cuadratica)

}

func main() {
	matriz := Matrix{alto: 4, ancho: 4}

	matriz.Set(2, 3, 4, 5, 6, 7, 8, 9, 20)
	matriz.Print()
}
