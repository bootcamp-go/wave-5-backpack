package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoInsercion(elementos []int, c chan float64) {
	tiempo := time.Now()
	for i := 1; i < len(elementos); i++ {
		for j := i - 1; j >= 0 && elementos[j] > elementos[i]; j-- {
			elementos[j+1] = elementos[j]
			elementos[j] = elementos[i]
		}
	}
	c <- float64(time.Since(tiempo))
}

func ordenamientoBurbuja(lista []int, c chan float64) {
	tiempo := time.Now()
	var value int
	for i := 0; i < len(lista); i++ {
		for j := 0; j < len(lista); j++ {
			if lista[i] > lista[j] {
				value = lista[i]
				lista[i] = lista[j]
				lista[j] = value
			}
		}
	}
	c <- float64(time.Since(tiempo))
}

func ordenamientoSeleccion(arreglo []int, c chan float64) {
	tiempo := time.Now()
	for i := 0; i < len(arreglo); i++ {
		min, min_pos := arreglo[i], i

		actual := arreglo[i]
		for j := i + 1; j < len(arreglo); j++ {
			valorActualList := arreglo[j]
			if valorActualList < min {
				min, min_pos = valorActualList, j
			}
		}

		if min != actual {
			arreglo[i], arreglo[min_pos] = min, actual
		}
	}
	c <- float64(time.Since(tiempo))
}

func imprimirResult(number int, valueInsercion, valuerBurbuja, valueSeleccion float64) {
	fmt.Printf("----------\nPara %d\n----------\n", number)
	fmt.Println("Tiempo ordenamiento con inserción:", valueInsercion)
	fmt.Println("Tiempo ordenamiento con burbuja:", valuerBurbuja)
	fmt.Println("Tiempo ordenamiento con selección:", valueSeleccion)
}

func main() {
	//Ejemplo con 100
	cien := rand.Perm(100)

	canalUnoC := make(chan float64)
	canalDosC := make(chan float64)
	canalTresC := make(chan float64)

	go ordenamientoInsercion(cien, canalUnoC)
	tiempoTotalInsercionC := <-canalUnoC

	go ordenamientoBurbuja(cien, canalDosC)
	tiempoTotalBurbujaC := <-canalDosC

	go ordenamientoSeleccion(cien, canalTresC)
	tiempoTotalSeleccionC := <-canalTresC

	imprimirResult(100, tiempoTotalInsercionC, tiempoTotalBurbujaC, tiempoTotalSeleccionC)

	//Ejemplo con 1000
	mil := rand.Perm(1000)

	canalUnoM := make(chan float64)
	canalDosM := make(chan float64)
	canalTresM := make(chan float64)

	go ordenamientoInsercion(mil, canalUnoM)
	tiempoTotalInsercionM := <-canalUnoM

	go ordenamientoBurbuja(mil, canalDosM)
	tiempoTotalBurbujaM := <-canalDosM

	go ordenamientoSeleccion(mil, canalTresM)
	tiempoTotalSeleccionM := <-canalTresM

	imprimirResult(1000, tiempoTotalInsercionM, tiempoTotalBurbujaM, tiempoTotalSeleccionM)

	//Ejemplo con 10000
	diez := rand.Perm(10000)

	canalUnoD := make(chan float64)
	canalDosD := make(chan float64)
	canalTresD := make(chan float64)

	go ordenamientoInsercion(diez, canalUnoD)
	tiempoTotalInsercionD := <-canalUnoD

	go ordenamientoBurbuja(diez, canalDosD)
	tiempoTotalBurbujaD := <-canalDosD

	go ordenamientoSeleccion(diez, canalTresD)
	tiempoTotalSeleccionD := <-canalTresD

	imprimirResult(10000, tiempoTotalInsercionD, tiempoTotalBurbujaD, tiempoTotalSeleccionD)
}
