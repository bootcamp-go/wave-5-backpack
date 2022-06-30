package main

import "fmt"

type Matrix struct {
	Valores    [][]float64
	Alto       int
	Ancho      int
	Cuadratica bool
	Maximo     float64
}

func (m *Matrix) Set(valores ...float64) {

	// Valido primer set
	if len(m.Valores) == 0 {
		m.Ancho = len(valores)
	}

	// Valido set posteriores
	if m.Ancho == len(valores) {
		m.Valores = append(m.Valores, valores)
		m.Alto = len(m.Valores)
		m.Cuadratica = m.Alto == m.Ancho
		m.setMax()
	}
}

func (m *Matrix) setMax() {
	max := 0.0
	for _, fila := range m.Valores {
		for _, columna := range fila {
			if max == 0.0 {
				max = columna
			}

			if columna > max {
				max = columna
			}
		}
	}
	m.Maximo = max
}

func (m Matrix) Print() {
	fmt.Println("Valores :")
	for _, fila := range m.Valores {
		for _, columna := range fila {
			fmt.Printf("%3.f  ", columna)
		}
		fmt.Println()
	}
	fmt.Println("Alto :", m.Alto)
	fmt.Println("Ancho :", m.Ancho)
	fmt.Println("Cuadratica :", m.Cuadratica)
	fmt.Println("Maximo :", m.Maximo)
}

func main() {

	m := Matrix{}
	m.Set(1, 2, 3)
	m.Set(1, 5, 3)
	m.Set(10, 5, 3)
	m.Print()

}
