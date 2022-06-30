package main

import "fmt"

type Matriz struct {
	valores     []float64
	alto, ancho int
	cuadratica  bool
	maximo      float64
}

func (m *Matriz) set(valores ...float64) {
	m.valores = valores
	max := valores[0]
	for _, valor := range valores {
		if valor > max {
			max = valor
		}
	}
	m.maximo = max
}

func (m Matriz) print() {
	pos := 0
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			if pos >= len(m.valores) {
				fmt.Printf("%f ", 0.0)
			} else {
				fmt.Printf("%f ", m.valores[pos])
			}
			pos++
		}
		fmt.Println()
	}
}

func main() {
	m1 := Matriz{alto: 3, ancho: 3}
	m1.set(1, 2, 3, 4)
	m1.print()
}
