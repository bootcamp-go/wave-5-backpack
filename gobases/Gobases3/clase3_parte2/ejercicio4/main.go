package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Info struct {
	Resultado []int
	Tiempo    time.Duration
}

func insertionSort(c chan Info, valores ...int) {
	tiempo := time.Now()
	var j, clave int

	for i := 1; i < len(valores); i++ {
		clave = valores[i]
		j = i - 1
		for j >= 0 && valores[j] > clave {
			valores[j+1] = valores[j]
			j -= 1
		}
		valores[j+1] = clave
	}

	info := &Info{Resultado: valores, Tiempo: time.Since(tiempo)}
	c <- *info
}

func ordenamientoBurbuja(c chan Info, valores ...int) {
	tiempo := time.Now()
	for i := 1; i < len(valores)-1; i++ {
		for j := i + 1; j < len(valores); j++ {
			if valores[i] > valores[j] {
				aux := valores[i]
				valores[i] = valores[j]
				valores[j] = aux
			}
		}
	}

	info := &Info{Resultado: valores, Tiempo: time.Since(tiempo)}
	c <- *info
}

func ordenamientoSeleccion(c chan Info, valores ...int) {
	tiempo := time.Now()
	var valorMenor, posMenor, aux int

	for i := 0; i < len(valores); i++ {
		valorMenor = valores[i]
		posMenor = i

		for j := 0; j < len(valores); j++ {
			if valores[j] < valorMenor {
				valorMenor = valores[j]
				posMenor = j
			}
		}

		if posMenor != i {
			aux = valores[i]
			valores[i] = valorMenor
			valores[posMenor] = aux
		}
	}

	info := &Info{Resultado: valores, Tiempo: time.Since(tiempo)}
	c <- *info

}

func print(t1 *Info, t2 *Info, t3 *Info) {
	fmt.Println()
	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento burbuja de", t1.Tiempo)
	fmt.Println("Resultado burbuja: ", t1.Resultado)
	fmt.Println()

	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento inserción de", t2.Tiempo)
	fmt.Println("Resultado Inserción: ", t2.Resultado)
	fmt.Println()

	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento selección de", t3.Tiempo)
	fmt.Println("Resultado Selección: ", t3.Resultado)
	fmt.Println()
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c1 := make(chan Info)
	c2 := make(chan Info)
	c3 := make(chan Info)

	go insertionSort(c1, variable1...)
	go ordenamientoBurbuja(c2, variable1...)
	go ordenamientoSeleccion(c3, variable1...)

	t1Insertion := <-c1
	t1Burbuja := <-c2
	t1Seleccion := <-c3

	print(&t1Insertion, &t1Burbuja, &t1Seleccion)

	c4 := make(chan Info)
	c5 := make(chan Info)
	c6 := make(chan Info)

	go insertionSort(c4, variable2...)
	go ordenamientoBurbuja(c5, variable2...)
	go ordenamientoSeleccion(c6, variable2...)

	t2Insertion := <-c4
	t2Burbuja := <-c5
	t2Seleccion := <-c6

	print(&t2Insertion, &t2Burbuja, &t2Seleccion)

	c7 := make(chan Info)
	c8 := make(chan Info)
	c9 := make(chan Info)

	go insertionSort(c7, variable3...)
	go ordenamientoBurbuja(c8, variable3...)
	go ordenamientoSeleccion(c9, variable3...)

	t3Insertion := <-c7
	t3Burbuja := <-c8
	t3Seleccion := <-c9

	print(&t3Insertion, &t3Burbuja, &t3Seleccion)
}
