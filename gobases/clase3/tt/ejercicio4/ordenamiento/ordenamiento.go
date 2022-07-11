package ordenamiento

import (
	"fmt"
	"math"
	"time"
)

// Recibe un slice de int y los ordena con tres algoritmos distintos mediante gorutines.
// Estos algoritmos son Insertion, Bubble y Selection
func Ordenar(nums []int) {
	i, b, s := copy(nums), copy(nums), copy(nums)

	chanI, chanB, chanS := make(chan time.Duration), make(chan time.Duration), make(chan time.Duration)
	go Insertion(i, chanI)
	go Burbuja(b, chanB)
	go Seleccion(s, chanS)

	ordenados := make(map[string]time.Duration)

	ordenados["insertion"] = <-chanI
	ordenados["burbuja"] = <-chanB
	ordenados["seleccion"] = <-chanS

	medirTiempo(ordenados)
}

func copy(nums []int) []int {
	copy := make([]int, len(nums))

	for i, v := range nums {
		copy[i] = v
	}

	return copy
}

// Toma los tiempos de los algoritmos e imprime el de menor duración
func medirTiempo(maps map[string]time.Duration) {
	var min int64 = math.MaxInt64
	var minKey string

	for k, v := range maps {
		fmt.Printf("Algoritmo: %v. Tiempo: %v\n", k, v.Microseconds())

		if v.Microseconds() < min {
			min = v.Microseconds()
			minKey = k
		}
	}

	fmt.Printf("Algoritmo con menos tiempo: %v. Tiempo: %v\n\n", minKey, min)
}

// Ordena el slice de menor a mayor. Envia el tiempo de ejecución
func Insertion(nums []int, ch chan time.Duration) {
	inicio := time.Now()

	var aux int
	for i := 1; i < len(nums); i++ {
		aux = nums[i]
		for j := i - 1; j >= 0 && nums[j] > aux; j-- {
			nums[j+1] = nums[j]
			nums[j] = aux
		}
	}

	ch <- time.Since(inicio)
}

// Ordena el slice de menor a mayor. Envia el tiempo de ejecución
func Burbuja(nums []int, ch chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	ch <- time.Since(inicio)
}

// Ordena el slice de menor a mayor. Envia el tiempo de ejecución
func Seleccion(nums []int, ch chan time.Duration) {
	inicio := time.Now()

	for i := 0; i < len(nums); i++ {
		actual, minValue, minPos := nums[i], nums[i], i

		// encontrar el minimo en la parte desordenada
		for j := i + 1; j < len(nums); j++ {
			comparacion := nums[j]
			if comparacion < minValue {
				minValue, minPos = comparacion, j
			}
		}

		if minValue != actual {
			nums[i], nums[minPos] = minValue, actual
		}
	}

	ch <- time.Since(inicio)
}
