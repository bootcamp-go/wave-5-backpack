/*
	Ejercicio 5 - Calcular cantidad de alimento

	Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. 
	Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan 
	haber muchos más animales que refugiar.

	- perro necesitan 10 kg de alimento
	- gato 5 kg
	- Hamster 250 gramos.
	- Tarántula 150 gramos.

	Se solicita:
	Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal 
	especificado y que retorne una función y un mensaje (en caso que no exista el animal)
	Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del 
	tipo de animal especificado.
*/

package main

import "fmt"

const (
	perro = "perro"
	gato = "gato"
	tarantula = "tarántula"
	hamster = "hamster"
)

func main() {
	
}

func animal(animal string) (float64) {
	var resultado float64
	switch animal {
	case perro:
		return calculaAlimento(valores)
	case gato:
		return calculoMinimo(valores)
	case tarantula:
		return calculoMaximo(valores)
	}
case hamster :
default:
		return "El animal no existe"
	return resultado, nil
}

// PENDIENTE