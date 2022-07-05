package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(array []int, c chan int) {
	now := time.Now()
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	last := time.Now()
	diff := last.Sub(now)
	fmt.Println("Bubble: ", diff.Minutes())
	c <- 1
}

func insertionsort(items []int, c chan int) {
	now := time.Now()
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
	last := time.Now()
	diff := last.Sub(now)
	fmt.Println("Insertion: ", diff.Minutes())
	c <- 2
}

func selectionsort(items []int, c chan int) {
	now := time.Now()
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
	last := time.Now()
	diff := last.Sub(now)
	fmt.Println("Selection: ", diff.Minutes())
	c <- 3
}

func main() {
	v1 := rand.Perm(100)
	v2 := rand.Perm(1000)
	v3 := rand.Perm(100000)
	c := make(chan int)
	go BubbleSort(v1, c)
	go insertionsort(v1, c)
	go selectionsort(v1, c)

	for i := 0; i < 3; i++ {
		<-c
	}
	fmt.Println()
	go BubbleSort(v2, c)
	go insertionsort(v2, c)
	go selectionsort(v2, c)

	for i := 0; i < 3; i++ {
		<-c
	}
	fmt.Println()
	go BubbleSort(v3, c)
	go insertionsort(v3, c)
	go selectionsort(v3, c)

	for i := 0; i < 3; i++ {
		<-c
	}
}
