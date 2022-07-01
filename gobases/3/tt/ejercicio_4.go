package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Resultado struct {
	ListaOrdenada []int
	Tiempo        time.Duration
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c1 := make(chan Resultado)
	c2 := make(chan Resultado)
	c3 := make(chan Resultado)

	go Insercion(variable1, c1)
	go Burbuja(variable1, c2)
	go Seleccion(variable1, c3)

	t1 := <-c1
	t2 := <-c2
	t3 := <-c3

	Imprimir(t1, t2, t3, "variable1")

	c4 := make(chan Resultado)
	c5 := make(chan Resultado)
	c6 := make(chan Resultado)

	go Insercion(variable2, c4)
	go Burbuja(variable2, c5)
	go Seleccion(variable2, c6)

	t4 := <-c4
	t5 := <-c5
	t6 := <-c6

	Imprimir(t4, t5, t6, "variable2")

	c7 := make(chan Resultado)
	c8 := make(chan Resultado)
	c9 := make(chan Resultado)

	go Insercion(variable3, c7)
	go Burbuja(variable3, c8)
	go Seleccion(variable3, c9)

	t7 := <-c7
	t8 := <-c8
	t9 := <-c9

	Imprimir(t7, t8, t9, "variable3")
}

func Imprimir(i Resultado, b Resultado, s Resultado, variable string) {
	fmt.Printf("### %s ###\n", variable)

	fmt.Printf("- Tiempo Inserción: %v\n\n", i.Tiempo)
	//fmt.Printf("%v\n\n", i.ListaOrdenada)

	fmt.Printf("- Tiempo Burbuja: %v\n\n", b.Tiempo)
	//fmt.Printf("%v\n\n", b.ListaOrdenada)

	fmt.Printf("- Tiempo Seleccion: %v\n\n", s.Tiempo)
	//fmt.Printf("%v\n\n", s.ListaOrdenada)
}

func Insercion(ListaDesordenada []int, c chan Resultado) {
	start := time.Now()
	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	Info := &Resultado{ListaOrdenada: ListaDesordenada, Tiempo: time.Since(start)}
	c <- *Info
}

func Burbuja(ListaDesordenada []int, c chan Resultado) {
	start := time.Now()
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		for j := 0; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[i] > ListaDesordenada[j] {
				auxiliar = ListaDesordenada[i]
				ListaDesordenada[i] = ListaDesordenada[j]
				ListaDesordenada[j] = auxiliar
			}
		}
	}
	Info := &Resultado{ListaOrdenada: ListaDesordenada, Tiempo: time.Since(start)}
	c <- *Info
}

func Seleccion(ListaDesordenada []int, c chan Resultado) {
	start := time.Now()
	var auxiliar int
	for i := 0; i < len(ListaDesordenada); i++ {
		min := i
		for j := i + 1; j < len(ListaDesordenada); j++ {
			if ListaDesordenada[j] < ListaDesordenada[min] {
				min = j
			}
		}
		if i != min {
			auxiliar = ListaDesordenada[i]
			ListaDesordenada[i] = ListaDesordenada[min]
			ListaDesordenada[min] = auxiliar
		}
	}
	Info := &Resultado{ListaOrdenada: ListaDesordenada, Tiempo: time.Since(start)}
	c <- *Info
}
