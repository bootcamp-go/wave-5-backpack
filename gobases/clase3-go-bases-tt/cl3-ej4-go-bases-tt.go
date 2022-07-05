/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #4:  Ordenamiento
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		It is required to instantiate 3 arrays with unordered random values
			- an array of integers with 100 values
			- an array of integers with 1000 values
			- an array of integers with 10000 values

		Para instanciar las variables utilizar rand
			package  main
			import  (
				"math/rand"
			)
			func  main  () {
			variable1  := rand.  Perm  (  100  )
			variable2  := rand.  Perm  (  1000  )
			variable3  := rand.  Perm  (  10000  )
		Each one must be sorted by:
			- Sorting by insertion
			- Sorting by bubble
			- Sorting by selection
		A go routine for each sorting run.
		I must wait for the 100 numbers to finish the sorting to follow
		the 1000 and then the 10000.
		Finally I must measure the time of each one and show on screen
		the result, to know which sorting was better for each array.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//	STRUCT : canal
type canal struct {
	nombre       string
	tiempoMedido time.Duration
}

//	FUNCTIONS
func insertionSort(lista []int, c chan canal) {
	start := time.Now()
	n := len(lista)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if lista[j-1] > lista[j] {
				lista[j-1], lista[j] = lista[j], lista[j-1]
			}
			j = j - 1
		}
	}
	c <- canal{"Ordenamiento por insercion", time.Since(start)}
}

func bubbleSort(lista []int, c chan canal) {
	start := time.Now()
	for i := 0; i < len(lista)-1; i++ {
		for j := 0; j < len(lista)-i-1; j++ {
			if lista[j] > lista[j+1] {
				lista[j], lista[j+1] = lista[j+1], lista[j]
			}
		}
	}
	c <- canal{"Ordenamiento por burbuja", time.Since(start)}
}

func selectionSort(lista []int, c chan canal) {
	start := time.Now()
	len := len(lista)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if lista[j] < lista[minIndex] {
				lista[j], lista[minIndex] = lista[minIndex], lista[j]
			}
		}
	}
	c <- canal{"Ordenamiento por seleccion", time.Since(start)}
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t\t|| Ordenamiento ||")
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c := make(chan canal)
	var tmp canal

	/* ----- Variable 1 ----- */
	go insertionSort(variable1, c)
	go bubbleSort(variable1, c)
	go selectionSort(variable1, c)

	fmt.Println("\n\t\t* Variable 1 | 100 elementos *")
	fmt.Println("\t  TIPO DE ORDENAMIENTO\t\tTIEMPO")
	fmt.Println("\t=========================================")
	for i := 0; i < 3; i++ {
		tmp = <-c
		fmt.Println("\t", tmp.nombre, "\t", tmp.tiempoMedido)
	}

	/* ----- Variable 2 ----- */
	go insertionSort(variable2, c)
	go bubbleSort(variable2, c)
	go selectionSort(variable2, c)

	fmt.Println("\n\t\t* Variable 2 | 1000 elementos *")
	fmt.Println("\t  TIPO DE ORDENAMIENTO\t\tTIEMPO")
	fmt.Println("\t=========================================")
	for i := 0; i < 3; i++ {
		tmp = <-c
		fmt.Println("\t", tmp.nombre, "\t", tmp.tiempoMedido)
	}

	/* ----- Variable 3 ----- */
	go insertionSort(variable3, c)
	go bubbleSort(variable3, c)
	go selectionSort(variable3, c)

	fmt.Println("\n\t\t* Variable 3 | 10000 elementos *")
	fmt.Println("\t  TIPO DE ORDENAMIENTO\t\tTIEMPO")
	fmt.Println("\t=========================================")
	for i := 0; i < 3; i++ {
		tmp = <-c
		fmt.Println("\t", tmp.nombre, "\t", tmp.tiempoMedido)
	}
}
