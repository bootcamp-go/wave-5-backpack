package main

import (
	"math/rand"
	"time"
)

func bubbleSort(values []int, c chan time.Duration) {
	start := time.Now()
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				temp := values[j]
				values[j] = values[j+1]
				values[j+1] = temp
			}
		}
	}
	difference := time.Since(start)
	c <- difference
}

func insertionSort(values []int, c chan time.Duration) {
	start := time.Now()
	for i := 1; i < len(values); i++ {
		key := values[i]
		j := i - 1

		for j >= 0 && values[j] > key {
			values[j+1] = values[j]
			j = j - 1
		}
		values[j+1] = key
	}
	difference := time.Since(start)
	c <- difference
}

func selectionSort(values []int, c chan time.Duration) {
	start := time.Now()
	for i := 0; i < len(values)-1; i++ {
		min_idx := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[min_idx] {
				min_idx = j
			}
		}
		temp := values[min_idx]
		values[min_idx] = values[i]
		values[i] = temp
	}
	difference := time.Since(start)
	c <- difference
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c1 := make(chan time.Duration)
	go bubbleSort(variable1, c1)
	go insertionSort(variable1, c1)
	go selectionSort(variable1, c1)

	for i := 0; i < 3; i++ {
		println(<-c1)
	}

	c2 := make(chan time.Duration)
	go bubbleSort(variable2, c2)
	go insertionSort(variable2, c2)
	go selectionSort(variable2, c2)

	for i := 0; i < 3; i++ {
		println(<-c2)
	}

	c3 := make(chan time.Duration)
	go bubbleSort(variable3, c3)
	go insertionSort(variable3, c3)
	go selectionSort(variable3, c3)

	for i := 0; i < 3; i++ {
		println(<-c3)
	}
}
