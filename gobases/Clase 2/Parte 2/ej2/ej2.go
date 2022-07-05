package main

import (
	"fmt"
)

type Matriz struct {
	valor      [][]float64
	x          int
	y          int
	cuadratica bool
	maximo     float64
}

func (m *Matriz) Set(x int, y int, params ...float64) {
	m.x = x
	m.y = y
	counter := 0
	mayor := params[0]
	for i := 0; i < m.y; i++ {
		m.valor = append(m.valor, []float64{})
		for j := 0; j < m.x; j++ {
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

func (m *Matriz) Print() {
	for i := 0; i < m.y; i++ {
		for j := 0; j < m.x; j++ {
			fmt.Printf("%.2f ", m.valor[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	mx := Matriz{}
	mx.Set(
		5, 7,
		1, 2, 3, 4, 5,
		6, 7, 8, 9, 10,
		11, 12, 13, 14, 15,
		16, 17, 18, 19, 20,
		21, 22, 23, 24, 25,
		26, 27, 28, 29, 30,
		31, 32, 33, 34, 35,
	)
	mx.Print()
	fmt.Println(mx.cuadratica, mx.maximo)
}
