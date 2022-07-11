package main

import "fmt"

type Matrix struct {
	alto int
	ancho int
	matriz [3][3]float64
	valueMax float64
}

func (m *Matrix) Set(values ...float64) {
	m.alto = 3
	m.ancho = 3
	aux := 0
	for i := 0; i < m.ancho; i++ {
		for j := 0; j < m.alto; j++ {
			m.matriz[i][j] = values[aux]
			aux++
		}
		
	}
}

func (m Matrix) Print()  {
	for i := 0; i < m.alto; i++ {
		fmt.Println(m.matriz[i])
	}
}

func main()  {
	matrix := Matrix{}
	matrix.Set(1,2,3,4,5,6,7,8,9)
	matrix.Print()
}