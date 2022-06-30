package main

import "fmt"

func main() {
	var m Matrix
	m.ancho = 3
	m.alto = 2
	m.Set(1, 2, 4, 5, 3, 5)
	m.Print()
}

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	cuadratica bool
	valMax     int
}

func (m *Matrix) Set(valores ...float64) {
	for _, valor := range valores {
		m.valores = append(m.valores, valor)
	}
}

func (m Matrix) Print() {
	aux := 0
	for i := 0; i < m.ancho; i++ {
		for j := 0; j < m.alto; j++ {
			fmt.Printf("%v\t", m.valores[aux])
			aux++
		}
		fmt.Println()
	}
}
