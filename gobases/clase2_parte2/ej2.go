/*

Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener 
a) los valores de la matriz, 
b) la dimensión del alto, 
c) la dimensión del ancho, 
d) si es cuadrática y 
e) cuál es el valor máximo.


*/

package main


import (
	"fmt"
)

type Matrix struct {
Valores 	 int
Alto 		float64
Ancho 		float64
Cuadratica 	bool
valorMaximo	float64
}

func main() {

	
	p1 := Persona {"Stefano", "Trejo", "36878354", "07/04/1992"}
	p2 := Persona {"Juan", "Gomez", "36123456", "07/04/1992"}
	p3 := Persona {"Pedro", "Basbus", "40456789", "07/04/1992"}
	
	p1.detalle()
	p2.detalle()
	p3.detalle()
}

func (persona Persona) detalle() {
	fmt.Println("Nombre: ", persona.Nombre)
	fmt.Println("Apellido: ", persona.Apellido)
	fmt.Println("DNI: ", persona.DNI)
	fmt.Println("Fecha: ", persona.Fecha)
}

