package main

import (
	// "errors"
	"fmt"
	"reflect"
)

type Matrix struct {
	valores         [][]int
	dimencion_alto  []int
	dimencion_ancho int
	cuadratica      bool
	valor_min       int
}

func allSameStrings(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
func (m *Matrix) set(valores_agre ...int) {
	m.dimencion_ancho = m.dimencion_ancho + 1
	m.dimencion_alto = append(m.dimencion_alto, len(valores_agre))
	var column []int
	for _, element := range valores_agre {
		column = append(column, element)
		if m.valor_min > element {
			m.valor_min = element
		}
	}

	if m.dimencion_ancho == m.dimencion_alto[0] && allSameStrings(m.dimencion_alto) {
		m.cuadratica = true
	} else {
		m.cuadratica = false
	}
	m.valores = append(m.valores, column)
}
func (m Matrix) PrintDetails() {
	fmt.Println("Person Details\n***************")
	p := reflect.ValueOf(m)
	for i := 0; i < p.NumField(); i++ {

		if i == 0 {
			fmt.Println("--------------")
			for _, columna := range m.valores {
				fmt.Println(columna)
			}
			fmt.Println("--------------")
		} else {
			fmt.Println(p.Field(i))
		}

	}
}
func main() {
	// polo := [][]int{{4, 3}, {5, 3}}
	// fmt.Println(polo)
	// polo = append(polo, []int{5, 8})
	// fmt.Println(polo)
	// matriz := Matrix{}
	// fmt.Println(matriz.valores)
	// matriz.set(4, 3, 5, 6)
	// fmt.Println(matriz.valores)
	// matriz.set(4, 2, 57, -6)
	// fmt.Println(matriz.valores)
	// matriz.PrintDetails()
	// fmt.Println(matriz.valores)
	// matriz.set(43, -5, 0)
	// fmt.Println(matriz.valores)
	// matriz.set(54, 32, -25, 12)
	// fmt.Println(matriz.valores)
	// matriz.PrintDetails()
}
