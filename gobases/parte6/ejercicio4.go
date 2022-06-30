package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordInsercion(numeros []int, ordenamiento chan string) {
	for i := 1; i < len(numeros); i++ {
		j := i
		numeroActual := numeros[j]
		for j > 0 && numeroActual < numeros[j-1] {
			numeros[j] = numeros[j-1]
			numeros[j-1] = numeroActual
			j--
		}
	}
	ordenamiento <- "Insercion"
}

func ordBurbuja(numeros []int, ordenamiento chan string) {
	for i := len(numeros) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if numeros[j] > numeros[j+1] {
				numeroActual := numeros[j]
				numeros[j] = numeros[j+1]
				numeros[j+1] = numeroActual
			}
		}
	}
	ordenamiento <- "Burbuja"
}

func ordSeleccion(numeros []int, ordenamiento chan string) {
	for i := 0; i < len(numeros)-1; i++ {
		min := numeros[i]
		pos := i
		for j := i + 1; j < len(numeros); j++ {
			if numeros[j] < min {
				min = numeros[j]
				pos = j
			}
		}
		if i != pos {
			numeros[pos] = numeros[i]
			numeros[i] = min
		}
	}
	ordenamiento <- "Seleccion"
}

func main() {

	ordenamiento := make(chan string)

	variables := [][]int{
		rand.Perm(100),
		rand.Perm(1000),
		rand.Perm(10000),
	}

	cantidadDeNumeros := 100

	for _, variable := range variables {

		tiempoInicial := time.Now()

		go ordInsercion(variable, ordenamiento)
		go ordBurbuja(variable, ordenamiento)
		go ordSeleccion(variable, ordenamiento)

		fmt.Printf("Para un array de %d valores:\n\n", cantidadDeNumeros)
		fmt.Printf("Puesto |   Algorimo   |    Tiempo   \n")
		fmt.Printf("------------------------------------\n")

		for i := 1; i <= 3; i++ {
			fmt.Printf("   %d   |  %10s  | %12s\n", i, <-ordenamiento, time.Since(tiempoInicial))
		}

		cantidadDeNumeros *= 10
		fmt.Printf("------------------------------------\n\n")
	}

}
