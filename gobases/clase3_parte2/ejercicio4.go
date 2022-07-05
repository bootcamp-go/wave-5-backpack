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

func ordIns(c chan Info, num []int) {
	start := time.Now()
	var aux int
	for i := 1; i < len(num); i++ {
		aux = num[i]
		for j := i - 1; j >= 0 && num[j] > aux; j-- {
			num[j+1] = num[j]
			num[j] = aux
		}
	}
	//return num
	info := &Info{Resultado: num, Tiempo: time.Since(start)}
	c <- *info
}

func ordBur(c chan Info, num []int) {
	start := time.Now()
	var aux int
	for i := 0; i < len(num); i++ {
		for j := 0; j < len(num); j++ {
			if num[i] > num[j] {
				aux = num[i]
				num[i] = num[j]
				num[j] = aux
			}
		}
	}
	info := &Info{Resultado: num, Tiempo: time.Since(start)}
	c <- *info
}

func ordSel(c chan Info, num []int) {
	start := time.Now()
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				var aux int
				aux = num[i]
				num[i] = num[j]
				num[j] = aux
			}
		}
	}
	info := &Info{Resultado: num, Tiempo: time.Since(start)}
	c <- *info
}

func Print(t1 *Info, t2 *Info, t3 *Info, elementos int) {
	fmt.Println()
	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento burbuja de", elementos, " 🕛 ", t1.Tiempo)
	//fmt.Println("Resultado burbuja: ", t1.Resultado)
	fmt.Println()

	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento inserción de", elementos, " 🕛 ", t2.Tiempo)
	//fmt.Println("Resultado Inserción: ", t2.Resultado)
	fmt.Println()

	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento selección de", elementos, "🕛 ", t3.Tiempo)
	//fmt.Println("Resultado Selección: ", t3.Resultado)
	fmt.Println()
}

func main() {
	//Cracion de Arrays
	var1 := rand.Perm(100)
	var2 := rand.Perm(1000)
	var3 := rand.Perm(10000)

	c1 := make(chan Info)
	c2 := make(chan Info)
	c3 := make(chan Info)

	go ordBur(c1, var1)
	go ordIns(c2, var1)
	go ordSel(c3, var1)

	rBur1 := <-c1
	rIns1 := <-c2
	rSel1 := <-c3

	Print(&rBur1, &rIns1, &rSel1, len(var1))

	c4 := make(chan Info)
	c5 := make(chan Info)
	c6 := make(chan Info)

	go ordBur(c4, var2)
	go ordIns(c5, var2)
	go ordSel(c6, var2)

	rBur2 := <-c4
	rIns2 := <-c5
	rSel2 := <-c6

	Print(&rBur2, &rIns2, &rSel2, len(var2))

	c7 := make(chan Info)
	c8 := make(chan Info)
	c9 := make(chan Info)

	go ordBur(c7, var3)
	go ordIns(c8, var3)
	go ordSel(c9, var3)

	rBur3 := <-c7
	rIns3 := <-c8
	rSel3 := <-c9

	Print(&rBur3, &rIns3, &rSel3, len(var3))
}
