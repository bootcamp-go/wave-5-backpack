/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #2:  Matrix
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		An artificial intelligence company needs to have a functionality to
		create a structure that represents a data matrix.
		For this they require a Matrix structure that has the methods:
			‣ Set: Receives a series of floating point values and initializes
			  the values in the Matrix structure.
			‣ Print: Prints the matrix on the screen in a more visible way
			  (with line breaks between rows).
		The Matrix structure must contain the values of the matrix, the height
		dimension, the width dimension, if it is quadratic and what is the
		maximum value.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import (
	"fmt"
)

//	STRUCT : Matrix
type Matrix struct {
	Ancho    int
	Alto     int
	Cuadrada bool
	Maximo   float64
	matriz   [][]float64
}

//	FUNCTION : (*Matrix).Set()
func (m *Matrix) Set(valAncho int, valAlto int, values ...float64) {
	m.Alto = valAlto
	m.Ancho = valAncho

	if m.Alto == m.Ancho { // Si es 'Cuadratica'
		m.Cuadrada = true
	} else {
		m.Cuadrada = false
	}

	if (m.Alto * m.Ancho) != len(values) { // Error en caso que los valores no entren en la matriz n x m
		fmt.Println("** La cantidad de valores no corresponde a la matriz n x m **")
		return
	}

	m.matriz = make([][]float64, valAlto)

	for i := 0; i < valAlto; i++ {
		m.matriz[i] = make([]float64, valAncho)
	}
	m.Maximo = values[0]
	for i := 0; i < valAlto; i++ {
		for j := 0; j < valAncho; j++ {
			m.matriz[i][j] = values[i*valAncho+j]
			if m.Maximo < values[i*valAncho+j] {
				m.Maximo = values[i*valAncho+j]
			}

		}
	}
}

//	STRUCT : (Matrix).Print()
func (m Matrix) Print() {
	fmt.Println("\n> Cuadratica: ", m.Cuadrada)
	fmt.Println("> Valor maximo: ", m.Maximo)
	fmt.Println("")
	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Print(m.matriz[i][j], "  ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Matrix ||")

	// Matriz 1 - Cuadratica
	var mtx Matrix
	mtx.Set(2, 2, 1.0, 2.0, 3.0, 4.0)
	mtx.Print()

	// Matriz 2 - No Cuadratica
	var mtx2 Matrix
	mtx2.Set(4, 2, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0)
	mtx2.Print()
}
