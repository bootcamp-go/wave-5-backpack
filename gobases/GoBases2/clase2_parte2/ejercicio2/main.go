package main

import "fmt"

type Matrix struct {
	valor      [][]float64
	x          int
	y          int
	cuadratica bool
	maximo     float64
}

func (m *Matrix) Set(x, y int, params ...float64) {
	m.x = x
	m.y = y
	counter := 0
	mayor := params[0]

	for i := 0; i < m.x; i++ {
		m.valor = append(m.valor, []float64{})

		for j := 0; j < m.y; j++ {
			m.valor[i] = append(m.valor[i], params[counter])

			if params[counter] > mayor {
				mayor = params[counter]
			}
			counter++
		}
	}

	m.maximo = mayor
	if m.x == m.y {
		m.cuadratica = true
	}
}

func (m *Matrix) Print() {
	for i := 0; i < m.x; i++ {
		for j := 0; j < m.y; j++ {
			fmt.Printf("%.2f ", m.valor[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	mx := Matrix{}

	mx.Set(2, 2, 1, 2, 3, 4)
	mx.Print()
	fmt.Println(mx.cuadratica, mx.maximo)
}
