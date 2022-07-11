package main

import "fmt"

type Matrix struct {
	matrix       [][]float64
	alto         int
	ancho        int
	isCuadratica bool
	valMax       float64
}

func (m *Matrix) Set(alto, ancho int, valores ...float64) {
	m.alto = alto
	m.ancho = ancho
	m.isCuadratica = (alto == ancho)
	var index int
	newMatriz := make([][]float64, alto)
	for i := 0; i < alto; i++ {
		newFile := make([]float64, ancho)
		for j := 0; j < ancho; j++ {
			if index < len(valores) {
				newFile[j] = valores[index]
				if valores[index] > m.valMax {
					m.valMax = valores[index]
				}
				index++
			}
		}
		newMatriz[i] = newFile
	}
	m.matrix = newMatriz
}

func (m *Matrix) Print() {
	fmt.Println("Imprimiendo la Matrix")
	for i := 0; i < m.alto; i++ {
		file := ""
		for j := 0; j < m.ancho; j++ {
			file += fmt.Sprintf("[%2.2f]  ", m.matrix[i][j])
		}
		fmt.Println(file)
	}
}

func main() {
	var matrix Matrix
	matrix.Set(3, 3, 1, 100, 5, 1, 6, 8)
	matrix.Print()
}
