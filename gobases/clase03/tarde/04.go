/*
Ejercicio 4 - Ordenamiento
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo


*/

package main
 
import (
   "math/rand"
   "fmt"
   "time"
)

type Info struct {
	Resultado []int
	Tiempo    time.Duration
}

func Insercion(c chan Info, valores ...int) {
	var (
		tiempo = time.Now()
		j, clave int
	)
	for i:=1; i < len(valores); i++ {
		clave = valores[i]
		j = i-1

		//esto sería como un while
		for j>=0 && valores[j]>clave {

			// inserto valor donde corresponde
			valores[j+1] = valores[j]

			j -= 1
		}

		valores[j+1] = clave
	}

	info := &Info{Resultado: valores, Tiempo: time.Since(tiempo)}
	c <- *info
}

func Burbuja(c chan Info, valores ...int) {
	var tiempo = time.Now()
	for i:=0; i < len(valores); i++ {
		for j:=i+1; j < len(valores); j++ {
			if valores[i]>valores[j]{
				aux := valores[i]
				valores[i] = valores[j]
				valores[j] = aux
			}
		}
	}

	info := &Info{Resultado: valores, Tiempo: time.Since(tiempo)}
	c <- *info
}

func Seleccion(c chan Info, valores ...int) {
	tiempo := time.Now()
	var min, indexMin, aux int
	for i := 0; i < len(valores); i++ {
		min = valores[i]
		indexMin = i

		for j := i + 1; j < len(valores); j++ {
			if valores[j] < min {
				min = valores[j]
				indexMin = j
			}
		}

		if indexMin != i {
			aux = valores[i]
			valores[i] = min
			valores[indexMin] = aux
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

	go Burbuja(c1, variable1...)
	go Insercion(c2, variable1...)
	go Seleccion(c3, variable1...)

	tBurbuja := <-c1
	tInsercion := <-c2
	tSeleccion := <-c3

	Print(&tBurbuja, &tInsercion, &tSeleccion, len(variable1))

	c4 := make(chan Info)
	c5 := make(chan Info)
	c6 := make(chan Info)

	go Burbuja(c4, variable2...)
	go Insercion(c5, variable2...)
	go Seleccion(c6, variable2...)

	t2Burbuja := <-c4
	t2Insercion := <-c5
	t2Seleccion := <-c6

	Print(&t2Burbuja, &t2Insercion, &t2Seleccion, len(variable2))

	c7 := make(chan Info)
	c8 := make(chan Info)
	c9 := make(chan Info)

	go Burbuja(c7, variable3...)
	go Insercion(c8, variable3...)
	go Seleccion(c9, variable3...)

	t3Burbuja := <-c7
	t3Insercion := <-c8
	t3Seleccion := <-c9

	Print(&t3Burbuja, &t3Insercion, &t3Seleccion, len(variable3))
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


