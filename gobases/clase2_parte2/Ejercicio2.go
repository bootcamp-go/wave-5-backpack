package main

import "fmt"

type Matrix struct {
	values    []float64
	height    int
	width     int
	cuadratic bool
	maxValue  float64
}

func (m *Matrix) set(values []float64) {
	m.values = values
}
func (m Matrix) print() {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			fmt.Printf(m[i][j])
		}
	}

}
func main() {

}
