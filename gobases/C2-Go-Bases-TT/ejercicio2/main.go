package main

import "fmt"

// Ejercicio 2 - Matrix
// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear
// una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:

//  - Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
//  - Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)

// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho,
// si es cuadrática y cuál es el valor máximo.

type Matrix struct {
	Ancho    int
	Alto     int
	Cuadrada bool
	MaxValor float64
	Valores  []float64
}

func (m *Matrix) Set(valores ...float64) {
	// Creamos un arreglo con el tamaño de la lista
	m.Valores = make([]float64, len(valores))

	// Seteamos la lista en el arreglo
	for i, v := range valores {
		m.Valores[i] = v
	}

	// Verificamos si es divisible entre 2 o 3 para determinar el alto y ancho
	if len(m.Valores)%2 == 0 {
		m.Ancho = len(m.Valores) / 2
		m.Alto = 2
	} else if len(m.Valores)%3 == 0 {
		m.Ancho = len(m.Valores) / 3
		m.Alto = 3
	}

	// Validamos si la matriz es cuadrada o no
	if m.Ancho == m.Alto {
		m.Cuadrada = true
	} else {
		m.Cuadrada = false
	}

	// Iteramos los elementos del arreglo para obtener el mayor elemento
	firstElement := true
	for _, v := range m.Valores {
		if firstElement {
			// Obtenemos la primera calificación de referencia
			m.MaxValor = v
			firstElement = false
		} else if v > m.MaxValor {
			// Validamos las demás calificaciones
			m.MaxValor = v
		}
	}
}

func (m Matrix) Print() {
	// Mostramos si la matriz es cuadrada
	if m.Cuadrada {
		fmt.Printf("La matriz SI es cuadrada\n")
	} else {
		fmt.Printf("La matriz NO es cuadrada\n")
	}

	// Mostramos el valor máximo de la matriz
	fmt.Printf("El valor máximo de la matriz es: %.2f\n", m.MaxValor)

	// Mostramos la matriz
	for i := 0; i < m.Alto; i++ {
		for j := 0; j < m.Ancho; j++ {
			fmt.Printf("[%.2f]", m.Valores[i+j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("Ejercicio 2 - Matrix")
	fmt.Println("")

	// Creamos una matriz
	var mat Matrix
	mat.Set(2.0, 4.0, 3.0, 5.0, 4.0, 6.0, 5.0, 7.0, 6.0)
	mat.Print()
}
