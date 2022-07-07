package main

import (
	"fmt"
	"math/rand"
	"time"
)

type results struct {
	v    []int
	time time.Duration
}

func main() {

	cienNumeros := rand.Perm(100)
	milNumeros := rand.Perm(1000)
	diezMNumeros := rand.Perm(10000)

	c1 := make(chan results)
	c2 := make(chan results)
	c3 := make(chan results)

	// con cien numeros
	go insercion(c1, cienNumeros...)
	go seleccion(c2, cienNumeros...)
	go burbuja(c3, cienNumeros...)

	tInsercion := <-c1
	tSeleccion := <-c2
	tBurbuja := <-c3
	fmt.Printf("Arreglo con Insercion: %v \n con un tiempo de %v:\n\n", tInsercion.v[80:], &tInsercion.time)
	fmt.Printf("Arreglo con Seleccion: %v \n con un tiempo de %v:\n\n", tSeleccion.v[80:], tSeleccion.time)
	fmt.Printf("Arreglo con Burbuja: %v \n con un tiempo de %v:\n\n", tBurbuja.v[80:], tBurbuja.time)
	fmt.Println("*********************************************************************************************************")

	// con mil numeros
	c4 := make(chan results)
	c5 := make(chan results)
	c6 := make(chan results)
	go insercion(c4, milNumeros...)
	go seleccion(c5, milNumeros...)
	go burbuja(c6, milNumeros...)

	tInsercion2 := <-c4
	tSeleccion2 := <-c5
	tBurbuja2 := <-c6
	fmt.Printf("Arreglo con Insercion: %v ...\n con un tiempo de %v :\n\n", tInsercion2.v[980:], &tInsercion2.time)
	fmt.Printf("Arreglo con Seleccion: %v... \n con un tiempo de %v:\n\n", tSeleccion2.v[980:], tSeleccion2.time)
	fmt.Printf("Arreglo con Burbuja: %v... \n con un tiempo de %v:\n\n", tBurbuja2.v[980:], tBurbuja2.time)
	fmt.Println("*********************************************************************************************************")

	// con diez numeros
	c7 := make(chan results)
	c8 := make(chan results)
	c9 := make(chan results)
	go insercion(c7, diezMNumeros...)
	go seleccion(c8, diezMNumeros...)
	go burbuja(c9, diezMNumeros...)

	tInsercion3 := <-c7
	tSeleccion3 := <-c8
	tBurbuja3 := <-c9
	fmt.Printf("Arreglo con Insercion: %v ...\n con un tiempo de %v :\n\n", tInsercion3.v[9980:], &tInsercion3.time)
	fmt.Printf("Arreglo con Seleccion: %v... \n con un tiempo de %v:\n\n", tSeleccion3.v[9980:], tSeleccion3.time)
	fmt.Printf("Arreglo con Burbuja: %v... \n con un tiempo de %v:\n\n", tBurbuja3.v[9980:], tBurbuja3.time)
	fmt.Println("*********************************************************************************************************")

}

func insercion(c chan results, valores ...int) {
	tiempo := time.Now()

	var (
		aux int
		j   int
	)
	for i := 1; i < len(valores); i++ {
		aux = valores[i]
		j = i - 1
		for j >= 0 && valores[j] > aux {
			var auxJ int = valores[j]
			valores[j] = aux
			valores[j+1] = auxJ

			j -= 1
		}
	}

	res := &results{time: time.Since(tiempo), v: valores}
	c <- *res

}

func burbuja(c chan results, valores ...int) {
	tiempo := time.Now()
	for i := 1; i < len(valores); i++ {
		for j := 0; j < len(valores); j++ {
			if valores[j] > valores[i] {

				aux := valores[i]
				valores[i] = valores[j]
				valores[j] = aux

			}
		}

	}

	res := &results{time: time.Since(tiempo), v: valores}
	c <- *res
}

func seleccion(c chan results, valores ...int) {
	tiempo := time.Now()

	var auxJ, auxI int

	for i := 0; i < len(valores); i++ {
		miniumValue := valores[i]
		for j, nextValue := range valores[i+1:] {

			if nextValue < miniumValue {

				miniumValue = nextValue
				auxJ = j + (i + 1)
				auxI = valores[i]

				valores[i] = valores[auxJ]
				valores[auxJ] = auxI
			}

		}

	}

	res := &results{time: time.Since(tiempo), v: valores}
	c <- *res
}
