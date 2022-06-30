package main

import (
	"fmt"
	"math"
)
//Ejercicio 2 - Matrix

type Matrix struct {
	valores    []float64
	alto       int
	ancho      int
	cuadratica bool
	max_val    int
}

func (m *Matrix) Set(valores []float64) {
	m.valores = valores
}

func (m Matrix) Print() {
	var pos int = 0
	for i := 0; i < m.alto; i++ {
		var fila string = ""
		for j := 0; j < m.ancho; j++ {
			if pos < len(m.valores) {
				val := fmt.Sprintf("%.1f", m.valores[pos])
				fila += val + "\t"
			} else {
				fila += fmt.Sprintf("®") + "\t"
			}
			pos++
		}
		fmt.Println(fila)
	}
}

func valMax(valores []float64) float64 {
	var res float64 = 0.0
	for i, value := range valores {
		if i == 0 {
			res = value
		} else {
			res = math.Max(res, value)
		}
	}
	return res
}

func main() {
	m1 := Matrix{alto: 3, ancho: 4, cuadratica: true, max_val: 5}
	var valores = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m1.Set(valores)
	m1.Print()
	fmt.Println("----------------------------")
	fmt.Print("El maximo valor de la Matriz es: ", valMax(valores), 
"\n")
}

