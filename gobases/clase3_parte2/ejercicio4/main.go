package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const (
	INS = "insercion"
	BUR = "burbuja"
	SEL = "selecion"
)

func ordenar(tipoOrden string) (func(numbers []int, c chan int) []int, error) {
	switch tipoOrden {
	case INS:
		return ordenInsercion, nil
	case BUR:
		return ordenBurbuja, nil
	case SEL:
		return ordenSeleccion, nil
	default:
		return nil, errors.New("Tipo de orden no valido")
	}
}

func ordenInsercion(numbers []int, c chan int) []int {
	fmt.Printf("Comienza ordenamiento por insercion para %d elementos\n", len(numbers))
	start := time.Now()
	for i := 1; i < len(numbers); i++ {
		j := i
		for j > 0 {
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
			j = j - 1
		}
	}
	c <- int(time.Since(start).Microseconds())
	return numbers
}

func ordenBurbuja(numbers []int, c chan int) []int {
	fmt.Printf("Comienza ordenamiento por burbuja para %d elementos\n", len(numbers))
	start := time.Now()
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	c <- int(time.Since(start).Microseconds())
	return numbers
}

func ordenSeleccion(numbers []int, c chan int) []int {
	fmt.Printf("Comienza ordenamiento por seleccion para %d elementos\n", len(numbers))
	start := time.Now()
	for i := 0; i < len(numbers); i++ {
		var minIdx = i
		for j := i; j < len(numbers); j++ {
			if numbers[j] < numbers[minIdx] {
				minIdx = j
			}
		}
		numbers[i], numbers[minIdx] = numbers[minIdx], numbers[i]
	}
	c <- int(time.Since(start).Microseconds())
	return numbers
}

func medirTiempo(ordenType string) {
	fmt.Println("=================================")
	arreglo1 := rand.Perm(100)
	arreglo2 := rand.Perm(1000)
	arreglo3 := rand.Perm(10000)
	orden, err := ordenar(ordenType)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		total := 0
		c1 := make(chan int)
		c2 := make(chan int)
		c3 := make(chan int)
		go orden(arreglo1, c1)
		t1 := <-c1
		fmt.Println("Se demoro :", t1)
		go orden(arreglo2, c2)
		t2 := <-c2
		fmt.Println("Se demoro :", t2)
		go orden(arreglo3, c3)
		t3 := <-c3
		fmt.Println("Se demoro :", t3)
		total = t1 + t2 + t3
		fmt.Println("En total se demora :", total)
	}
}

func main() {
	medirTiempo(INS)
	medirTiempo(BUR)
	medirTiempo(SEL)
}
