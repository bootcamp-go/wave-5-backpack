/*Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento
fue mejor para cada arreglo
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	arregloCien := rand.Perm(100)
	//arregloMil := rand.Perm(1000)
	//arregloDiezMil := rand.Perm(10000)

	go OrdenInsercion(arregloCien, c1)
	tiempoTotalInsercion := <-c1
	fmt.Println("el tiempo que demoro el ArregloCien en ORden insercion fue,", tiempoTotalInsercion)

	go OrdenBurbuja(arregloCien, c2)
	tiempoTotalBurbuja := <-c2
	fmt.Println("el tiempo que demoro el ArregloCien en Orden burbuja fue,", tiempoTotalBurbuja)

	go OrdenSeleccion(arregloCien, c3)
	tiempoTotalSeleccion := <-c3
	fmt.Println("el tiempo que demoro el ArregloCien en Orden seleccion fue,", tiempoTotalSeleccion)

	//ACA DEBO HACER NUEVOS CANALES Y REPETIR EL MISMO PROCESO ANTERIOR CON LOS OTROS 2 ARREGLOS X CADA UNO DE LOS ORDENAMIENTOS

}

func OrdenInsercion(ListaDesordenada []int, c1 chan float64) {
	tiempoInicio := time.Now()
	var auxiliar int
	for i := 1; i < len(ListaDesordenada); i++ {
		auxiliar = ListaDesordenada[i]
		for j := i - 1; j >= 0 && ListaDesordenada[j] > auxiliar; j-- {
			ListaDesordenada[j+1] = ListaDesordenada[j]
			ListaDesordenada[j] = auxiliar
		}
	}
	elapsed := float64(time.Since(tiempoInicio))
	c1 <- elapsed
}

func OrdenBurbuja(ListaDesordenada []int, c2 chan float64) {
	tiempoInicio := time.Now()
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
	elapsed := time.Since(tiempoInicio)
	c2 <- float64(elapsed)
}

func OrdenSeleccion(arreglo []int, c3 chan float64) {
	tiempoInicio := time.Now()
	for i := 0; i < len(arreglo); i++ {
		minimo_encontrado, posicion_minimo := arreglo[i], i

		valor_original := arreglo[i]
		// encontrar minimo en parte desordenada
		for j := i + 1; j < len(arreglo); j++ {
			valor_comparacion := arreglo[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}

		if minimo_encontrado != valor_original {
			// intercambio posiciones con primer desordenado
			arreglo[i], arreglo[posicion_minimo] = minimo_encontrado, valor_original
		}
	}
	elapsed := time.Since(tiempoInicio)
	c3 <- float64(elapsed)
}
