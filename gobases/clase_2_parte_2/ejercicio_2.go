package main

import "fmt"

type Matrix struct {
	matriz     [][]float32
	maximo     float32
	alto       int
	ancho      int
	cuadratica bool
}

func (m *Matrix) Set(puntos ...float32) {

	m.matriz = append(m.matriz, puntos)

	ancho := len(puntos)

	var valorMaximo float32
	for i, punto := range puntos {
		if i == 0 {
			valorMaximo = punto
		}
		if punto > valorMaximo {
			valorMaximo = punto
		}
	}

	if m.alto == 0 {
		m.maximo = valorMaximo
	}

	if valorMaximo > m.maximo {
		m.maximo = valorMaximo
	}

	if m.ancho == 0 {
		m.ancho = ancho
		m.cuadratica = true
	}

	if m.ancho != ancho {
		m.ancho = ancho
		m.cuadratica = false
	}

	m.alto += 1
	m.cuadratica = m.alto == m.ancho
}

func (m *Matrix) Print() {
	for _, line := range m.matriz {
		for _, number := range line {
			fmt.Printf("%3.2f  ", number)

		}
		fmt.Println()
	}
	fmt.Println("Alto", m.alto)
	fmt.Println("Ancho", m.ancho)
	fmt.Println("Es Cuadratica", m.cuadratica)
	fmt.Println("Valor Máximo", m.maximo)
}

func main() {
	matrix := Matrix{}
	matrix.Set(14.5, 12, 23, 22.5)
	matrix.Set(9, 1, 30, 21.5)
	matrix.Set(12.5, 11, 23, 26.5)
	matrix.Set(12.5, 11, 23, 26.5)

	matrix.Print()
}
