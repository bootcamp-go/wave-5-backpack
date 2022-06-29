package main

import "fmt"

//Estructura matriz "Matrix"
type Matrix struct {
	valor      [][]float64
	x          int
	y          int
	cuadratica bool
	maximo     float64
}

//Inicializando los valores de la matriz en Set
func (m *Matrix) Set(x int, y int, params ...float64) {
	m.x = y
	m.y = x
	counter := 0
	mayor := params[0]
	for i := 0; i < m.x; i++ {
		m.valor = append(m.valor, []float64{})
		for j := 0; j < m.y; j++ {
			m.valor[i] = append(m.valor[i], params[counter])
			//Para saber el mayor
			if params[counter] > mayor {
				mayor = params[counter]
			}
			counter++
		}
	}
	m.maximo = mayor
	//Tomando por cuadrática que tenga ancho y alto iguales
	if m.x == m.y {
		m.cuadratica = true
	}
}

//Imprimiendo la matriz
func (m *Matrix) Print() {
	for i := 0; i < m.x; i++ {
		for j := 0; j < m.y; j++ {
			fmt.Printf("%.2f ", m.valor[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	mx := Matrix{}
	fmt.Println("Matrix")
	mx.Set(3, 5, 1, 2, 3, 4, 5, 6,7,8,9,10,11,12,13,14,15)
	mx.Print()
	fmt.Println("¿Es cuadrática?",mx.cuadratica,"\nNúmero máximo",mx.maximo)
}