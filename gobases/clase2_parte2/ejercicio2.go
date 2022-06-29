package main

import "fmt"

type Matrix struct {
	values   []float64
	alto     int
	ancho    int
	cuadrada bool
	max_val  float64
}

func main() {
	matrix1 := Matrix{
		alto:     3,
		ancho:    3,
		cuadrada: true,
		max_val:  5,
	}
	var values = []float64{1, 2, 3, 4, 5}
	matrix1.set(values)
	matrix1.get()
}

func (matrix *Matrix) set(values []float64) {
	matrix.values = values
}

func (matrix Matrix) get() {
	var pos int = 0
	for i := 0; i < matrix.alto; i++ {
		var fila string = ""
		for j := 0; j < matrix.ancho; j++ {
			if pos < len(matrix.values) {
				val := fmt.Sprintf("%f", matrix.values[pos])
				fila += val + "\t"
			} else {
				fila += fmt.Sprintf("%f", 0.0) + "\t"
			}
			pos++
		}
		fmt.Println(fila)
	}
}
