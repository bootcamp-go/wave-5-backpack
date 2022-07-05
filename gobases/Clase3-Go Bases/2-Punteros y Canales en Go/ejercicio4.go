package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variable1 := rand.Perm(100)
	//fmt.Printf("%v \n", variable1)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	cb := make(chan time.Duration)
	cI := make(chan time.Duration)
	cS := make(chan time.Duration)
	go ordenamientoBurbuja(&variable1, cb)
	go ordenamientoInsercion(&variable1, cI)
	go ordenamientoSeleccion(&variable1, cS)
	resultado := <-cb
	fmt.Printf("Burbuja %v \n", resultado)
	resultado = <-cI
	fmt.Printf("Insercion %v \n", resultado)
	resultado = <-cS
	fmt.Printf("Seleccion %v \n", resultado)
	go ordenamientoBurbuja(&variable2, cb)
	go ordenamientoInsercion(&variable2, cI)
	go ordenamientoSeleccion(&variable2, cS)
	resultado = <-cb
	fmt.Printf("Burbuja %v \n", resultado)
	resultado = <-cI
	fmt.Printf("Insercion %v \n", resultado)
	resultado = <-cS
	fmt.Printf("Seleccion %v \n", resultado)
	go ordenamientoBurbuja(&variable3, cb)
	go ordenamientoInsercion(&variable3, cI)
	go ordenamientoSeleccion(&variable3, cS)
	resultado = <-cb
	fmt.Printf("Burbuja %v \n", resultado)
	resultado = <-cI
	fmt.Printf("Insercion %v \n", resultado)
	resultado = <-cS
	fmt.Printf("Seleccion %v \n", resultado)

}

func ordenamientoInsercion(listaInt *[]int, c chan time.Duration) {
	inicio := time.Now()
	var auxiliar int
	for i := 1; i < len(*listaInt); i++ {
		auxiliar = (*listaInt)[i]
		for j := i - 1; j >= 0 && (*listaInt)[j] > auxiliar; j-- {
			(*listaInt)[j+1] = (*listaInt)[j]
			(*listaInt)[j] = auxiliar
		}
	}
	fmt.Println("Finalizó Insercion")
	final := time.Now()
	c <- (final.Sub(inicio))
}

func ordenamientoBurbuja(listaInt *[]int, c chan time.Duration) {
	inicio := time.Now()
	var auxiliar int
	for i := 0; i < len(*listaInt); i++ {
		for j := 0; j < len(*listaInt); j++ {
			if (*listaInt)[i] < (*listaInt)[j] {
				auxiliar = (*listaInt)[i]
				(*listaInt)[i] = (*listaInt)[j]
				(*listaInt)[j] = auxiliar
			}
		}
	}
	fmt.Println("Finalizó Burbuja")
	final := time.Now()
	c <- (final.Sub(inicio))
}

func ordenamientoSeleccion(listaInt *[]int, c chan time.Duration) {
	inicio := time.Now()
	stepCounter := 1
	for i := 0; i < len((*listaInt))-1; i++ {
		menor := i
		for j := i + 1; j < len((*listaInt)); j++ {
			stepCounter++
			if (*listaInt)[menor] > (*listaInt)[j] {
				menor = j
			}
		}

		v := (*listaInt)[i]
		(*listaInt)[i] = (*listaInt)[menor]
		(*listaInt)[menor] = v
	}
	fmt.Println("Finalizó Seleccion")
	final := time.Now()
	c <- (final.Sub(inicio))
}

/*Ejercicio 4 - Ordenamiento

Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
)


func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)
}

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado,
para saber qué ordenamiento fue mejor para cada arreglo


*/
