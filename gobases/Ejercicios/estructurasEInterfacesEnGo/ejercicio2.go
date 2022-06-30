package main

import "fmt"

type Matriz struct {
	Alto        int
	Ancho       int
	Valores     []float64
	Cuadratica  bool
	ValorMaximo float64
}

func (matriz *Matriz) setMatriz(valores ...float64) {
	matriz.Valores = valores
}

func (matriz Matriz) printMatriz() {
	indice := 0
	for i := 0; i < matriz.Alto; i++ {
		for j := 0; j < matriz.Ancho; j++ {
			if indice < len(matriz.Valores) {
				fmt.Print(matriz.Valores[indice], " ")
			}
			indice++
		}
		fmt.Println()
	}
}
