package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array1, array2, array3 := randomArrays()
	c := make(chan time.Duration)
	ordenamiento100(c, array1)
	ordenamiento100(c, array2)
	ordenamiento100(c, array3)
}

func randomArrays() ([]int, []int, []int) {
	array1 := rand.Perm(100)
	array2 := rand.Perm(1000)
	array3 := rand.Perm(10000)
	return array1, array2, array3
}

func ordenamientoInsercion(array []int, c chan time.Duration) {
	start := time.Now()
	var n = len(array)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}
	c <- time.Since(start)
}

func ordenamientoBurbuja(array []int, c chan time.Duration) {
	start := time.Now()
	var (
		n      = len(array)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if array[i] > array[i+1] {
				array[i+1], array[i] = array[i], array[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	c <- time.Since(start)
}

func ordenamientoSeleccion(array []int, c chan time.Duration) {
	start := time.Now()
	var n = len(array)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}
	c <- time.Since(start)
}

func ordenamiento100(c chan time.Duration, array []int) {
	elementos := len(array)
	go ordenamientoInsercion(array, c)
	fmt.Printf("El ordenamiento de insercion para %d elementos tardó: %s\n", elementos, <-c)
	go ordenamientoBurbuja(array, c)
	fmt.Printf("El ordenamiento de burbuja para %d elementos tardó: %s\n", elementos, <-c)
	go ordenamientoSeleccion(array, c)
	fmt.Printf("El ordenamiento de seleccion para %d elementos tardó: %s\n\n", elementos, <-c)
}
