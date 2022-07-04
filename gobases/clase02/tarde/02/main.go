/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.
*/

package main
import (
	"fmt"
	"errors"
)

type Matrix struct {
	alto int
	ancho int
	cuadratica bool
	maxValue int
	valores []int
}

func (m *Matrix) set(valores []int) error{
	if len(valores) > m.alto * m.ancho {
		err := errors.New("La cantidad de valores ingresados es mayor a la capacidad de la matriz elegida")
		fmt.Print(err,"\n")
		return err
	}

	m.valores = valores
	max := valores[0]
	for _, valor := range valores {
		if valor > max{
			max = valor
		}
	}
	m.maxValue = max
	return nil
}

func (m Matrix) print() {
	if len(m.valores) == 0 {
		fmt.Print("La matriz está vacía \n")
		return
	}
	
	for i := 0 ; i < (m.alto*m.ancho); i++ {
		if i % m.ancho == 0 {
			fmt.Print("\n")
		}
		if i < len(m.valores) {
			fmt.Print(m.valores[i],"\t")
		} else {
			fmt.Print("-\t")
		}
	}
	fmt.Print("\n")
}

func main() {

	var m1 Matrix
	m1.alto = 3
	m1.ancho = 3

	m1.set([]int{1, 2, 3, 11111, 65, 76, 87, 87, 99999})
	m1.print()
}
