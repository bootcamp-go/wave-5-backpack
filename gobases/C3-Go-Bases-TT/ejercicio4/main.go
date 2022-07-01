package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Ejercicio 4 - Ordenamiento

// Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
// Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
//   - un arreglo de números enteros con 100 valores
//   - un arreglo de números enteros con 1000 valores
//   - un arreglo de números enteros con 10000 valores

// Para instanciar las variables utilizar rand
//   - variable := rand.Perm(100)

// Se debe realizar el ordenamiento de cada una por:
//   - Ordenamiento por inserción
//   - Ordenamiento por burbuja
//   - Ordenamiento por selección

// Una go routine por cada ejecución de ordenamiento.

// Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.

// Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor
// para cada arreglo

// Algoritmo de ordenamiento por inserción
func ordenamientoInsercion(array []int, c chan time.Time) {
	var validar int
	for i := 1; i < len(array); i++ {
		// Se obtiene el segundo elemento del arreglo
		validar = array[i]
		// Se intercalan los valores si el valor de la derecha es mayor que el valor de la izquierda
		for j := i - 1; j >= 0 && array[j] > validar; j-- {
			array[j+1] = array[j]
			array[j] = validar
		}
	}
	c <- time.Now()
}

// Algoritmo de ordenamiento por burbuja
func ordenamientoBurbuja(array []int, c chan time.Time) {
	var validar int
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			// Cada elemento del arreglo se valida con todos los elementos del arreglo
			if array[i] > array[j] {
				validar = array[i]
				array[i] = array[j]
				array[j] = validar
			}
		}
	}
	c <- time.Now()
}

// Algoritmo de ordenamiento por selección
func ordenamientoSeleccion(array []int, c chan time.Time) {
	for i := 0; i < len(array); i++ {
		pos := i
		min := array[i]

		respaldo := array[i]
		// Se valida cada elemento del arreglo para obtener el mínimo
		for j := i + 1; j < len(array); j++ {
			comparar := array[j]
			if comparar < min {
				min = comparar
				pos = j
			}
		}

		if min != respaldo {
			// Se intercambia de posición si los valores seleccionados son diferentes
			array[i] = min
			array[pos] = respaldo
		}
	}
	c <- time.Now()
}

func calculoInsercion(array1 []int, array2 []int, array3 []int) {
	// Creamos los canales para cada goroutine
	channel1 := make(chan time.Time)
	channel2 := make(chan time.Time)
	channel3 := make(chan time.Time)

	// Creamos las goroutines para el ordenamiento por inserción
	startTime := time.Now()
	go ordenamientoInsercion(array1, channel1)
	go ordenamientoInsercion(array2, channel2)
	go ordenamientoInsercion(array3, channel3)

	// Obtenemos los resultados del ordenamiento por inserción
	tiempoFinal1 := <-channel1
	tiempoFinal2 := <-channel2
	tiempoFinal3 := <-channel3

	// Se imprimen los resultados del ordenamiento por inserción
	fmt.Println("\nAlgoritmo de ordenamiento por inserción:")
	fmt.Println("\t- 100 elementos: ", tiempoFinal1.Sub(startTime))
	fmt.Println("\t- 1000 elementos: ", tiempoFinal2.Sub(startTime))
	fmt.Println("\t- 10000 elementos: ", tiempoFinal3.Sub(startTime))
}

func calculoBurbuja(array1 []int, array2 []int, array3 []int) {
	// Creamos los canales para cada goroutine
	channel1 := make(chan time.Time)
	channel2 := make(chan time.Time)
	channel3 := make(chan time.Time)

	// Creamos las goroutines para el ordenamiento por burbuja
	startTime := time.Now()
	go ordenamientoBurbuja(array1, channel1)
	go ordenamientoBurbuja(array2, channel2)
	go ordenamientoBurbuja(array3, channel3)

	// Obtenemos los resultados del ordenamiento por burbuja
	tiempoFinal1 := <-channel1
	tiempoFinal2 := <-channel2
	tiempoFinal3 := <-channel3

	// Se imprimen los resultados del ordenamiento por burbuja
	fmt.Println("\nAlgoritmo de ordenamiento por burbuja:")
	fmt.Println("\t- 100 elementos: ", tiempoFinal1.Sub(startTime))
	fmt.Println("\t- 1000 elementos: ", tiempoFinal2.Sub(startTime))
	fmt.Println("\t- 10000 elementos: ", tiempoFinal3.Sub(startTime))
}

func calculoSeleccion(array1 []int, array2 []int, array3 []int) {
	// Creamos los canales para cada goroutine
	channel1 := make(chan time.Time)
	channel2 := make(chan time.Time)
	channel3 := make(chan time.Time)

	// Creamos las goroutines para el ordenamiento por selección
	startTime := time.Now()
	go ordenamientoSeleccion(array1, channel1)
	go ordenamientoSeleccion(array2, channel2)
	go ordenamientoSeleccion(array3, channel3)

	// Obtenemos los resultados del ordenamiento por burbuja
	tiempoFinal1 := <-channel1
	tiempoFinal2 := <-channel2
	tiempoFinal3 := <-channel3

	// Se imprimen los resultados del ordenamiento por burbuja
	fmt.Println("\nAlgoritmo de ordenamiento por selección:")
	fmt.Println("\t- 100 elementos: ", tiempoFinal1.Sub(startTime))
	fmt.Println("\t- 1000 elementos: ", tiempoFinal2.Sub(startTime))
	fmt.Println("\t- 10000 elementos: ", tiempoFinal3.Sub(startTime))
}

func main() {
	fmt.Println("Ejercicio 4 - Ordenamiento")

	// Creamos las variables que serán ordenadas
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	// Validamos los algoritmos
	calculoInsercion(variable1, variable2, variable3)
	calculoBurbuja(variable1, variable2, variable3)
	calculoSeleccion(variable1, variable2, variable3)
}
