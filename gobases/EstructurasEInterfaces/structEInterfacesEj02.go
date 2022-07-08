package main

// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.

// Para ello requieren una estructura Matrix que tenga los métodos:

// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix

// Print: Imprime por pantalla la matriz de una forma más visible (Con los saltos de línea entre filas)

// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

type Matrix struct {
	valores      [2][4]float64
	alto         int
	ancho        int
	esCuadratica bool
	maxValor     float64
}

func (m *Matrix) Set(valores ...float64) {

}

func main() {

}
