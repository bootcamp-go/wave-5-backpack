package main

import "fmt"

type Matrix struct {
	alto         int
	ancho        int
	isCuadratica bool
	valueMax     float64
	values       [3][3]float64
}

func main() {
	matriz := Matrix{}
	matriz.alto = 3
	matriz.ancho = 3
	matriz.isCuadratica = (matriz.alto == matriz.ancho)
	matriz.valueMax = 9
	// matriz.values = [matriz.alto][matriz.ancho]float64
	// fmt.Println(matriz)
	matriz.Set(1, 2, 3, 4, 5, 6, 7, 8, 9)
	matriz.Print()
}

func (m *Matrix) Set(values ...float64) {
	var aux int = 0
	for i := 0; i < m.ancho; i++ {
		for j := 0; j < m.alto; j++ {
			m.values[i][j] = values[aux]
			aux++
		}
	}
}

func (m Matrix) Print() {
	for i := 0; i < m.alto; i++ {
		fmt.Println(m.values[i])
	}
}
