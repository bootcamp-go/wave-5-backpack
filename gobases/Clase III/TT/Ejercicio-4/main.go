package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	result []int
	tiempo time.Duration
}

func insercion(c chan Data, valores ...int) {
	timeD := time.Now()
	for i, value := range valores {
		for j := i - 1; j >= 0; j-- {
			if value < valores[j] {
				valores[j+1] = valores[j]
				valores[j] = value
			}
		}
	}
	resultado := Data{result: valores, tiempo: time.Since(timeD)}

	c <- resultado
}

func burbuja(c chan Data, valores ...int) {
	timeD := time.Now()

	for j := 0; j <= len(valores)-1; j++ {
		for x := j + 1; x <= len(valores)-1; x++ {
			if valores[j] > valores[x] {
				valor_aux := valores[j]
				valores[j] = valores[x]
				valores[x] = valor_aux
			}
		}
	}
	resultado := Data{result: valores, tiempo: time.Since(timeD)}

	c <- resultado
}

func seleccion(c chan Data, valores ...int) {
	timeD := time.Now()
	var vmenor, min, aux int
	for j := 0; j < len(valores); j++ {
		min = j
		vmenor = valores[j]

		for x := j + 1; x < len(valores); x++ {
			if valores[x] < vmenor {
				vmenor = valores[x]
				min = x
			}
		}
		if min != j {
			aux = valores[j]
			valores[j] = vmenor
			valores[min] = aux
		}
	}

	resultado := Data{result: valores, tiempo: time.Since(timeD)}

	c <- resultado
}

func printResult(variable int, data1 *Data, data2 *Data, data3 *Data) {

	fmt.Println("-------------------------------------------")
	fmt.Println("El tiempo de función inserción para la variable con longitud ", variable, " es de: ", data1.tiempo)
	fmt.Println("-------------------------------------------")
	fmt.Println("El tiempo de función burbuja para la variable con longitud ", variable, " es de: ", data2.tiempo)
	fmt.Println("-------------------------------------------")
	fmt.Println("El tiempo de función selección para la variable con longitud ", variable, " es de: ", data2.tiempo)
	fmt.Println("-------------------------------------------")

}

func main() {

	variable1 := rand.Perm(10)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	// Cálculo tiempo con variable 1

	c1 := make(chan Data)
	c2 := make(chan Data)
	c3 := make(chan Data)

	go insercion(c1, variable1...)
	go burbuja(c2, variable1...)
	go seleccion(c3, variable1...)

	data1 := <-c1
	data2 := <-c2
	data3 := <-c3

	printResult(len(variable1), &data1, &data2, &data3)

	// Cálculo tiempo con variable 2

	c4 := make(chan Data)
	c5 := make(chan Data)
	c6 := make(chan Data)

	go insercion(c4, variable2...)
	go burbuja(c5, variable2...)
	go seleccion(c6, variable2...)

	data4 := <-c4
	data5 := <-c5
	data6 := <-c6

	printResult(len(variable2), &data4, &data5, &data6)

	// Cálculo tiempo con variable 3

	c7 := make(chan Data)
	c8 := make(chan Data)
	c9 := make(chan Data)

	go insercion(c7, variable3...)
	go burbuja(c8, variable3...)
	go seleccion(c9, variable3...)

	data7 := <-c7
	data8 := <-c8
	data9 := <-c9

	printResult(len(variable3), &data7, &data8, &data9)

}
