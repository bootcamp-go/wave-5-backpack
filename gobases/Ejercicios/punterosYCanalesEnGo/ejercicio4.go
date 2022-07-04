package main

import (
	"time"
)

func ordenamientoInsercion(c chan time.Duration, valores []int) {
	tiempo := time.Now()
	var j, clave int

	for i := 1; i < len(valores); i++ {
		clave = valores[i]
		j = i - 1

		for j >= 0 && valores[j] > clave {

			valores[j+1] = valores[j]
			j -= 1

		}
		valores[j+1] = clave
	}

	//fmt.Printf("Array ordenado", valores)
	c <- time.Since(tiempo)
}

func ordenamientoBurbuja(c chan time.Duration, valores []int) {
	tiempo := time.Now()
	for i := 0; i < len(valores)-1; i++ {
		for j := i + 1; j < len(valores); j++ {

			if valores[i] > valores[j] {
				aux := valores[i]
				valores[i] = valores[j]
				valores[j] = aux
			}
		}
	}

	c <- time.Since(tiempo)
}

func ordenamientoSeleccion(c chan time.Duration, valores []int) {
	tiempo := time.Now()
	var valorMenor, posMenor, aux int
	for i := 0; i < len(valores); i++ {
		valorMenor = valores[i]
		posMenor = i

		for j := i + 1; j < len(valores); j++ {
			if valores[j] < valorMenor {
				valorMenor = valores[j]
				posMenor = j
			}
		}

		if posMenor != i {
			aux = valores[i]
			valores[i] = valorMenor
			valores[posMenor] = aux
		}
	}

	c <- time.Since(tiempo)
}
