package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Resultado struct {
	nombre string
	tiempo int64
}

func InsertionSort(arr []int, c chan Resultado) {
	tiempo := time.Now()
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	c <- Resultado{"Order por inserción", time.Since(tiempo).Nanoseconds()}
}

func BubbleSort(arr []int, c chan Resultado) {
	tiempo := time.Now()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	c <- Resultado{"Order por burbuja", time.Since(tiempo).Nanoseconds()}
}

func SelectionSort(arr []int, c chan Resultado) {
	tiempo := time.Now()
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	c <- Resultado{"Ordenamiento por selección", time.Since(tiempo).Nanoseconds()}
}

func main() {
	arreglo1 := rand.Perm(100)
	arreglo2 := rand.Perm(1000)
	arreglo3 := rand.Perm(10000)

	c := make(chan Resultado)
	var resultadoTmp Resultado

	go InsertionSort(arreglo1, c)
	go BubbleSort(arreglo1, c)
	go SelectionSort(arreglo1, c)

	fmt.Println("Resultados para 100 elementos:")
	fmt.Println("Posición\t Nombre\t\t\t\t Tiempo")
	for i := 0; i < 3; i++ {
		resultadoTmp = <-c
		fmt.Println(i+1, "\t\t", resultadoTmp.nombre, "\t\t", resultadoTmp.tiempo, "ns")
	}
	fmt.Print("\n")

	go InsertionSort(arreglo2, c)
	go BubbleSort(arreglo2, c)
	go SelectionSort(arreglo2, c)

	fmt.Println("Resultados para 1000 elementos:")
	fmt.Println("Posición\t Nombre\t\t\t\t Tiempo")
	for i := 0; i < 3; i++ {
		resultadoTmp = <-c
		fmt.Println(i+1, "\t\t", resultadoTmp.nombre, "\t\t", resultadoTmp.tiempo, "ns")
	}
	fmt.Print("\n")

	go InsertionSort(arreglo3, c)
	go BubbleSort(arreglo3, c)
	go SelectionSort(arreglo3, c)

	fmt.Println("Resultados para 10000 elementos:")
	fmt.Println("Posición\t Nombre\t\t\t\t Tiempo")
	for i := 0; i < 3; i++ {
		resultadoTmp = <-c
		fmt.Println(i+1, "\t\t", resultadoTmp.nombre, "\t\t", resultadoTmp.tiempo, "ns")
	}
	fmt.Print("\n")

}
