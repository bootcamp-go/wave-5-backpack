package main

import "fmt"

type Matrix struct {
	Alto       int
	Ancho      int
	Max        float64
	Cuadratica bool
}

//Llenando el array
func (m *Matrix) set(alto, ancho int, valores ...float64) {
	m.Alto = alto
	m.Ancho = ancho
	const Matrix := [m.Alto][m.Ancho]float64{}
	for r = 0; r <= m.Alto; r++ {
		append(Matrix, 1)
		for c = 0; c < m.Ancho; c++ {
			append(Matrix, 2)
		}
	}
}

func (m set) print() {
	fmt.Printf(Matrix)
}

func main() {
	mx := Matrix{}
	mx.set(3, 3, 10, 11, 12)
	mx.print()
}
