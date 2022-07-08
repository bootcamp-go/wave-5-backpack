package main

import "fmt"

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	esCuadrada bool
	maxValue   float64
}

func (m *Matrix) set(values ...float64) {
	m.valores = values
}

func (m Matrix) print() string {
	pos := 0
	matrix := ""
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			matrix += fmt.Sprintf("%.2f\t", m.valores[pos])
			pos++
		}
		matrix += "\n"
	}

	return matrix
}

func main() {
	miMatriz := Matrix{
		alto:       2,
		ancho:      2,
		esCuadrada: true,
		maxValue:   1.0}

	miMatriz.set(1, 2, 3, 4)
	fmt.Print(miMatriz.print())
}

//Ejercicio realizado con los compañeros del meets
/*
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
*/
