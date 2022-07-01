package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*Ejercicio 4 - Ordenamiento

Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
	1. un arreglo de números enteros con 100 valores
	2. un arreglo de números enteros con 1000 valores
	3. un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
Se debe realizar el ordenamiento de cada una por:
	1. Ordenamiento por inserción
	2. Ordenamiento por burbuja
	3. Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento
fue mejor para cada arreglo.
*/

type Info struct {
	Resultado []int
	Tiempo    time.Duration
}

func ordenamientoInsercion(c chan Info, valores ...int) {
	tiempo := time.Now()
	var j, clave int
	for i := 1; i < len(valores); i++ {
		clave = valores[i]
		j = i - 1

		//fmt.Println("Comparando pares. . . (", valores[j], " - 🔑", clave, ")")
		for j >= 0 && valores[j] > clave {
			//time.Sleep(time.Second)
			//fmt.Println(valores[j], " > ", clave)
			// valores[j+1] = clave
			valores[j+1] = valores[j]
			//fmt.Println("Actualización", valores)
			j -= 1

		}

		valores[j+1] = clave
	}

	info := &Info{Resultado: valores, Tiempo: time.Since(tiempo)}
	c <- *info
}

func ordenamientoBurbuja(c chan Info, valores ...int) {
	tiempo := time.Now()
	for i := 0; i < len(valores)-1; i++ {
		for j := i + 1; j < len(valores); j++ {
			//time.Sleep(time.Second)
			if valores[i] > valores[j] {
				//fmt.Println("Comparando . . . ", valores[i], " > ", valores[j], "= SI")
				//fmt.Println("******************************")

				//time.Sleep(time.Second)
				//fmt.Println("Ordanmiento actual\t", valores)
				aux := valores[i]
				valores[i] = valores[j]
				valores[j] = aux
				//fmt.Println("¡Nuevo Ordenamiento!\t", valores)
				//fmt.Println("******************************")
			}
			//fmt.Println("Comparando . . . ", valores[i], " > ", valores[j], "= NO")
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

		for j := i + 1; j < len(valores); j++ {
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

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c1 := make(chan Info)
	c2 := make(chan Info)
	c3 := make(chan Info)

	go ordenamientoBurbuja(c1, variable1...)
	go ordenamientoInsercion(c2, variable1...)
	go ordenamientoSeleccion(c3, variable1...)

	tBurbuja := <-c1
	tInsercion := <-c2
	tSeleccion := <-c3

	Print(&tBurbuja, &tInsercion, &tSeleccion, len(variable1))

	c4 := make(chan Info)
	c5 := make(chan Info)
	c6 := make(chan Info)

	go ordenamientoBurbuja(c4, variable2...)
	go ordenamientoInsercion(c5, variable2...)
	go ordenamientoSeleccion(c6, variable2...)

	t2Burbuja := <-c4
	t2Insercion := <-c5
	t2Seleccion := <-c6

	Print(&t2Burbuja, &t2Insercion, &t2Seleccion, len(variable2))

	c7 := make(chan Info)
	c8 := make(chan Info)
	c9 := make(chan Info)

	go ordenamientoBurbuja(c7, variable3...)
	go ordenamientoInsercion(c8, variable3...)
	go ordenamientoSeleccion(c9, variable3...)

	t3Burbuja := <-c7
	t3Insercion := <-c8
	t3Seleccion := <-c9

	Print(&t3Burbuja, &t3Insercion, &t3Seleccion, len(variable3))
}

func Print(t1 *Info, t2 *Info, t3 *Info, elementos int) {
	fmt.Println()
	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento burbuja de", elementos, " 🕛 ", t1.Tiempo)
	fmt.Println("Resultado burbuja: ", t1.Resultado)
	fmt.Println()

	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento inserción de", elementos, " 🕛 ", t2.Tiempo)
	fmt.Println("Resultado Inserción: ", t2.Resultado)
	fmt.Println()

	fmt.Println("******************************************************************")
	fmt.Println("Tiempo de ejecución ordenamiento selección de", elementos, "🕛 ", t3.Tiempo)
	fmt.Println("Resultado Selección: ", t3.Resultado)
	fmt.Println()
}
