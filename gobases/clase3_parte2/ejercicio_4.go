package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionsort(items []int, c chan time.Duration) {
	start := time.Now()
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	c <- time.Since(start)
}

func bubblesort(items []int, c chan time.Duration) {
	start := time.Now()
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
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

func selectionsort(items []int, c chan time.Duration) {
	start := time.Now()
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	c <- time.Since(start)
}

func goRoutine(c chan time.Duration, arrays ...[]int) {

	for _, array := range arrays {
		go insertionsort(array, c)
		fmt.Printf("Para %d elementos el método Insertion sort se demora: %10s\n", len(array), <-c)

		go bubblesort(array, c)
		fmt.Printf("Para %d elementos el método Bubble sort se demora: %10s\n", len(array), <-c)

		go selectionsort(array, c)
		fmt.Printf("Para %d elementos el método Selection sort se demora: %10s\n\n", len(array), <-c)
	}
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c := make(chan time.Duration)

	goRoutine(c, variable1, variable2, variable3)
}
