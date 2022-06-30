package main

import "fmt"

type Matrix struct {
	valor      [][]float64
	x          int
	y          int
	cuadratica bool
	maximo     float64
}

func (m *Matrix) SetMatrix(x int, y int, params ...float64) {
	m.x = y
	m.y = x
	c := 0
	mayor := params[0]
	for i := 0; i < m.x; i++ {
		m.valor = append(m.valor, []float64{})
		for j := 0; j < m.y; j++ {
			m.valor[i] = append(m.valor[i], params[c])
			//Para saber el mayor
			if params[c] > mayor {
				mayor = params[c]
			}
			c++
		}
	}
	m.maximo = mayor
	if m.x == m.y {
		m.cuadratica = true
	} else {
		m.cuadratica = false
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
	fmt.Println("Matrix")
	mx.SetMatrix(8, 3, 2, 2, 5, 6, 7, 8, 9, 10, 11, 20)
	mx.Print()
	fmt.Println("¿Es cuadrática?", mx.cuadratica, "\nNúmero máximo", mx.maximo)
}
