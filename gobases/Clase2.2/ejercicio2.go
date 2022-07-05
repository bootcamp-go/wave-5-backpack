package main

import (
	"fmt"
)

type matrixStr struct {
	width, length int
	square        bool
	matrix        [][]float64
}

func (m matrixStr) printMatrix() {
	for _, value := range m.matrix {
		fmt.Println(value)
	}
}

func (m matrixStr) setMatrix(values ...float64) {
	i := 0
	for j := 0; j < m.width; j++ {
		for k := 0; k < m.length; k++ {
			if i >= len(values) {
				break
			}
			m.matrix[j][k] = values[i]
			i++
		}
	}
}

func newMatrix(width, length int) *matrixStr {
	m := matrixStr{width: width, length: length}
	m.matrix = make([][]float64, width)
	for i := range m.matrix {
		m.matrix[i] = make([]float64, length)
	}
	if width == length {
		m.square = true
	}
	return &m
}

func main() {
	m := newMatrix(4, 3)
	m.printMatrix()
	m.setMatrix(1, 2, 3, 4, 5, 6)
	m.printMatrix()
}
