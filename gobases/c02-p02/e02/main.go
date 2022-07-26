package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	matrix   [][]float64
	nFilas   int
	nCol     int
	cuadrada bool
	max      float64
}

func (m *Matrix) Set(nFilas, nCol int, elementos ...float64) error {

	if len(elementos) != nFilas*nCol {
		return errors.New("La cantidad de elementos ingresados es incorrecta.")
	}

	m.nFilas = nFilas
	m.nCol = nCol
	m.cuadrada = nFilas == nCol
	m.matrix = make([][]float64, nFilas)

	for i := 0; i < nFilas; i++ {
		m.matrix[i] = make([]float64, nCol)
		for j := 0; j < nCol; j++ {
			m.matrix[i][j] = elementos[nCol*i+j]
		}
	}

	return nil

}

func (m Matrix) obtenerMaximo() float64 {

	maximo := m.matrix[0][0]

	for i := 0; i < m.nFilas; i++ {
		for j := 0; j < m.nCol; j++ {
			if maximo < m.matrix[i][j] {
				maximo = m.matrix[i][j]
			}
		}
	}

	return maximo

}

func (m Matrix) detalle() {
	for i := 0; i < m.nFilas; i++ {
		for j := 0; j < m.nCol; j++ {
			fmt.Print(m.matrix[i][j], " ")
		}
		fmt.Println("")
	}
}

func main() {

	var m01 Matrix
	var _ error = m01.Set(2, 2, 1, 2, 3, 4)
	m01.detalle()

	var m02 Matrix
	m02.Set(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	m02.detalle()

	var m03 Matrix
	m03.Set(1, 1, 1)
	m03.detalle()

	var m04 Matrix
	error04 := m04.Set(1, 1, 2, 2)
	if error04 != nil {
		fmt.Println(error04)
	}

}
