package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func insertionSort(wg *sync.WaitGroup, c chan time.Duration, list_int []int) {
	defer wg.Done()
	start := time.Now()
	for post := 1; post < len(list_int); post++ {
		key := list_int[post]
		j := post - 1
		//fmt.Println(j, key, list_int[j])
		for j >= 0 && key < list_int[j] {

			list_int[j+1] = list_int[j]
			//fmt.Println(list_int)
			j -= 1
		}
		list_int[j+1] = key
	}
	elapsed := time.Since(start)
	//fmt.Println(list_int)
	c <- elapsed
}

func bubleSort(wg *sync.WaitGroup, c chan time.Duration, listNum []int) {
	defer wg.Done()
	start := time.Now()
	n := len(listNum)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if listNum[j] > listNum[j+1] {
				listNum[j], listNum[j+1] = listNum[j+1], listNum[j]
			}
		}
	}
	elapsed := time.Since(start)
	c <- elapsed
}

func selectSort(wg *sync.WaitGroup, c chan time.Duration, listInt []int) {
	defer wg.Done()
	start := time.Now()
	n := len(listInt)
	for i := 0; i < n; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if listInt[min_idx] > listInt[j] {
				min_idx = j
			}
		}
		listInt[i], listInt[min_idx] = listInt[min_idx], listInt[i]
	}
	elapsed := time.Since(start)
	c <- elapsed
}

func main() {
	fmt.Println("hola")
	var wg sync.WaitGroup
	//list_int := []int{9, 2, 7, 3, 5}
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	test_list := [][]int{variable1, variable2, variable3}

	c := make(chan time.Duration)
	d := make(chan time.Duration)
	e := make(chan time.Duration)
	for post, elem := range test_list {
		wg.Add(3)
		//selectSort
		go bubleSort(&wg, c, elem)
		go selectSort(&wg, d, elem)
		go insertionSort(&wg, e, elem)

		fmt.Println("bubleSort ", <-c)
		fmt.Println("selectSort ", <-d)
		fmt.Println("insertionSort ", <-e)
		wg.Wait()
		fmt.Println("se acabo todo ", post)
	}
}
