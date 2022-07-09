package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	Valores [] float64
	Alto int
	Ancho int
}

func main()  {
	m := Matrix{
		Valores: []float64{1, 30, 3, 45, 54, 25, 76, 97, 7, 10, 30, 65},
		Alto:    3,
		Ancho:   4,
	}
	Matrix.Set(m)
	Matrix.Print(m)
	Matrix.Cuadratica(m)
}

func (m Matrix) Set() {
	if len(m.Valores) != m.Ancho*m.Alto {
		fmt.Println("La cantidad de valores no coindice con las dimensiones especificadas")
	}
}

func (m Matrix) Print() {
	if len(m.Valores) == 0 {
		fmt.Println("La matriz esta vacia")
	}
	for fila := 0; fila < m.Alto; fila++ {
		fmt.Printf("%.0f\n", m.Valores[fila*m.Ancho:fila*m.Ancho+m.Ancho])
	}
}

func (m Matrix) Cuadratica() bool {
	if (m.Alto == m.Ancho) && m.Alto != 0 {
		return true
	}
	return false
}

func (m Matrix) Max() float64 {
	max := -math.MaxFloat64
	for _, elemento := range m.Valores {
		if elemento > max {
			max = elemento
		}
	}
	return max
}
