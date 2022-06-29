package main

import "fmt"

func main() {
	m := Matrix{}
	m.Set(3, 3, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	m.Print()

	m2 := Matrix{}
	m2.Set(4, 2, 1.14, 3.23, 4, 6, 4.44, 5.34, 4, 5)
	m2.Print()
}

type Matrix struct {
	val        [][]float64
	alto       int
	ancho      int
	cuadratica bool
	max        float64
}

// Recibe alto y ancho de la matriz. Luego inserta los valores pasados
// Si la lista de nums es mayor al tamaño de la matriz no se insertaran los números sobrantes
func (m *Matrix) Set(alto, ancho int, nums ...float64) {
	// set dimensiones
	m.alto = alto
	m.ancho = ancho

	// check cuadratica
	if alto == ancho {
		m.cuadratica = true
	}

	// n es un contador para iterar sobre nums
	n := 0

	// For para llenar [][]float64
	for y := 0; y < alto; y++ {
		ejeX := make([]float64, ancho) // crea slice con su ancho
		m.val = append(m.val, ejeX)

		for x := 0; x < m.ancho; x++ {
			num := nums[n]
			ejeX[x] = num

			// check maximo
			if num > m.max {
				m.max = num
			}

			n++
		}
	}
}

func (m Matrix) Print() {
	for y := 0; y < m.alto; y++ {
		fmt.Println(m.val[y])
	}

	fmt.Printf("Alto: %v\nAncho: %v\nCuadratica: %v\n", m.alto, m.ancho, m.cuadratica)
}
