package main

import "fmt"

type Matrix struct {
	alto         int
	ancho        int
	isCuadratica bool
	matriz       [3][3]float64
	valueMax     float64
}

func main() {
	//alto, ancho:= 5,5
	matriz := Matrix{}
	matriz.set(1, 2, 3, 4, 5, 6, 7, 8, 9)
	matriz.print()
}

func (m *Matrix) set(values ...float64) {
	// 1, 2, 3, |
	// 4, 5, 6, |
	// 7, 8, 9
	m.alto = 3
	m.ancho = 3
	var aux int = 0
	for i := 0; i < m.ancho; i++ {
		for j := 0; j < m.alto; j++ {
			m.matriz[i][j] = values[aux]
			aux++
		}
	}

	//fmt.Println(m.matriz)
}

func (m Matrix) print() {
	for i := 0; i < m.alto; i++ {
		fmt.Println(m.matriz[i])
	}

}
