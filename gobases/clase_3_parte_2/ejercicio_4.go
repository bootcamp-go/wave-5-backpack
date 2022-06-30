package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insercion(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func burbuja(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func seleccion(input []int) []int {
	stepCounter := 1
	for i := 0; i < len(input)-1; i++ {
		menor := i
		for j := i + 1; j < len(input); j++ {
			stepCounter++
			if input[menor] > input[j] {
				menor = j
			}
		}
		v := input[i]
		input[i] = input[menor]
		input[menor] = v
	}
	return input
}

type testingFunc func([]int) []int

func bench(oper testingFunc, data []int, executionTime chan time.Duration) {
	start := time.Now()
	oper(data)
	executionTime <- time.Since(start)
}

func runBenchmark(name string, data []int) {
	inserChan, burbChan, selChan := make(chan time.Duration), make(chan time.Duration), make(chan time.Duration)

	go bench(burbuja, data, burbChan)
	go bench(insercion, data, inserChan)
	go bench(seleccion, data, selChan)
	fmt.Printf("Benchmark Corriendo %s \n", name)
	fmt.Printf("Tiempo Insercion  %s \n", <-inserChan)
	fmt.Printf("Tiempo Burbuja    %s \n", <-burbChan)
	fmt.Printf("Tiempo Seleccion  %s \n", <-selChan)
	fmt.Printf("Benchmark Completado %s \n", name)
	println()
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	println()
	runBenchmark("Test de 100 números aleatorios", variable1)
	runBenchmark("Test de 1000 números aleatorios", variable2)
	runBenchmark("Test de 10000 números aleatorios", variable3)
}
