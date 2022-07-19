package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	matriz       [][]float64
	alto         int
	ancho        int
	esCuadratica bool
	valorMax     int
}

func (m *Matrix) Set(alto int, ancho int, valoresFlotantes ...float64) error {

	if len(valoresFlotantes) > alto*ancho {
		return errors.New("El número de valores es mayor que el número de celdas")
	} else if len(valoresFlotantes) < alto*ancho {
		return errors.New("El número de valores es menor que el número de celdas")
	}

	m.alto = alto
	m.ancho = ancho
	m.esCuadratica = alto == ancho

	m.matriz = make([][]float64, alto)

	for i := 0; i < alto; i++ {
		m.matriz[i] = make([]float64, ancho)
	}

	for i := 0; i < alto; i++ {
		for j := 0; j < ancho; j++ {
			m.matriz[i][j] = valoresFlotantes[(ancho*i)+j]
		}
	}

	return nil
}

func (m Matrix) Print() {
	for i := 0; i < m.alto; i++ {
		for j := 0; j < m.ancho; j++ {
			fmt.Print(m.matriz[i][j], " ")
		}
		fmt.Println("")
	}
}

func main() {
	var matri Matrix
	err := matri.Set(3, 3, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)
	if err != nil {
		fmt.Println(err)
	} else {
		matri.Print()
	}

	var matri2 Matrix
	err = matri2.Set(3, 3, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0)
	if err != nil {
		fmt.Println(err)
	} else {
		matri2.Print()
	}

	var matri3 Matrix
	err = matri3.Set(3, 3, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	if err != nil {
		fmt.Println(err)
	} else {
		matri3.Print()
	}

}
