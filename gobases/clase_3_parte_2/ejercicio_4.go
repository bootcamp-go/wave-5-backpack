package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SortAlgorithm struct {
	name      string
	operation SortingOperation
}

type SortingOperation func([]int) []int

type Bench interface {
	run() BenchResult
}

type SortingBench struct {
	algo SortAlgorithm
	data []int
}

func (s SortingBench) run() BenchResult {
	start := time.Now()
	s.algo.operation(s.data)
	return BenchResult{name: s.algo.name, timeExecution: time.Since(start)}
}

type Benchmark interface {
	run()
}

type AsyncBenchmark struct {
	name   string
	benchs []Bench
}

type BenchResult struct {
	name          string
	timeExecution time.Duration
}

func (ab *AsyncBenchmark) runAsyncBench(bench Bench, resulChan chan BenchResult) {
	result := bench.run()
	resulChan <- result
}

func (ab *AsyncBenchmark) run() {
	channels := make([]chan BenchResult, 0)
	for _, bench := range ab.benchs {
		benchChannel := make(chan BenchResult)
		go ab.runAsyncBench(bench, benchChannel)
		channels = append(channels, benchChannel)

	}
	fmt.Printf("Async Benchmark Corriendo %s \n", ab.name)
	for _, channel := range channels {
		result := <-channel
		fmt.Printf("Tiempo para %s  %s \n", result.name, result.timeExecution)
	}
	fmt.Printf("Async Benchmark Termino %s \n", ab.name)
}

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

func bench(algo SortAlgorithm, data []int, executionTime chan time.Duration) {
	start := time.Now()
	algo.operation(data)
	executionTime <- time.Since(start)
}

func newSortingsToBenchs(sortingAlgos []SortAlgorithm, data []int) []Bench {
	benchs := make([]Bench, 0)
	for _, sortingAlgo := range sortingAlgos {
		benchs = append(benchs, SortingBench{
			algo: sortingAlgo,
			data: data,
		})
	}
	return benchs
}

type UseCase struct {
	name string
	data []int
}

func newAsyncBenchmarksSortingAlgorthms(sortingOpes []SortAlgorithm, tests ...UseCase) []AsyncBenchmark {
	benchMarks := make([]AsyncBenchmark, 0)
	for _, test := range tests {
		benchs := newSortingsToBenchs(sortingOpes, test.data)
		benchMarks = append(benchMarks, AsyncBenchmark{
			name:   test.name,
			benchs: benchs,
		})
	}
	return benchMarks
}

func main() {
	sortingsAlgorthms := []SortAlgorithm{
		{name: "insercion", operation: insercion},
		{name: "burbuja", operation: burbuja},
		{name: "seleccion", operation: seleccion}}
	benchMarks := newAsyncBenchmarksSortingAlgorthms(
		sortingsAlgorthms,
		UseCase{name: "100 datos", data: rand.Perm(100)},
		UseCase{name: "1000 datos", data: rand.Perm(1000)},
		UseCase{name: "10000 datos", data: rand.Perm(10000)},
	)

	for _, benchMark := range benchMarks {
		benchMark.run()
		fmt.Println()
	}

}
