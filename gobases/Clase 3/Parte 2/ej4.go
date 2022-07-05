package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func InsertionSort(lista []int) []int {
	for i := 1; i < len(lista); i++ {
		j := i
		for j > 0 {
			if lista[j-1] > lista[j] {
				lista[j-1], lista[j] = lista[j], lista[j-1]
			}
			j--
		}
	}
	return lista
}

func SelectionSort(lista []int) []int {
	for i := 0; i < len(lista); i++ {
		var minIndex = i
		for j := i; j < len(lista); j++ {
			if lista[j] < lista[minIndex] {
				minIndex = j
			}
		}
		lista[i], lista[minIndex] = lista[minIndex], lista[i]
	}
	return lista
}

type funcionOrdenamiento func([]int) []int

func medirTiempoEjecucion(funcion funcionOrdenamiento, lista []int, chDuracion chan time.Duration) {
	start := time.Now()
	funcion(lista)
	diff := time.Since(start)
	chDuracion <- diff
}

func medirTiempoLista(lista []int) {
	lonLista := len(lista)
	chDuration1 := make(chan time.Duration)
	chDuration2 := make(chan time.Duration)
	chDuration3 := make(chan time.Duration)
	go medirTiempoEjecucion(BubbleSort, lista, chDuration1)
	go medirTiempoEjecucion(InsertionSort, lista, chDuration2)
	go medirTiempoEjecucion(SelectionSort, lista, chDuration3)
	fmt.Printf("Bubble Sort %d %v\n", lonLista, <-chDuration1)
	fmt.Printf("Insertion Sort %d %v\n", lonLista, <-chDuration2)
	fmt.Printf("Selection Sort %d %v\n", lonLista, <-chDuration3)
	fmt.Println()
}

func main() {
	lista1 := rand.Perm(100)
	lista2 := rand.Perm(1000)
	lista3 := rand.Perm(10000)
	medirTiempoLista(lista1)
	medirTiempoLista(lista2)
	medirTiempoLista(lista3)
}
