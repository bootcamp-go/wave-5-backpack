// Ejercicio 5 - Calcular cantidad de alimento

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

// perro necesitan 10 kg de alimento
// gato 5 kg
// Hamster 250 gramos.
// Tarántula 150 gramos.

// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

package main

import "fmt"


const (
	dog ="dogo"
	cat = "kitty"
	hamster = "hamster"
	tarantula = "spidey"
)
func dogFunc(dogs int) float64 {
	return 10 * float64(dogs)
}
func catFunc(cats int) float64 {
	return 5 * float64(cats)
}
func hamFunc(hamsters int) float64 {
	return 0.250 * float64(hamsters)
}
func tarFunc(tarantulas int) float64 {
	return 0.150 * float64(tarantulas)
}


func animal(animalType string) (func(int) float64, string ){
	switch animalType {
	case dog:
		return dogFunc, ""
	case cat:
		return catFunc, ""
	case hamster:
		return hamFunc, ""
	case tarantula:
		return tarFunc, ""
	default:
		return nil, "El animal no existe en la lista"
	}
}
func main() {
	var total float64

	dogFunc, msg := animal(dog)
	if msg != "" {
		fmt.Println(msg)
	} else {
		total += dogFunc(2) // 20 kgs
	}

	catFunc, msg := animal(cat)
	if msg != "" {
		fmt.Println(msg)
	} else {
		total += catFunc(5) 
	}

	hamFunc, msg := animal(hamster)
	if msg != "" {
		fmt.Println(msg)
	} else {
		total += hamFunc(4)
	}

	tarFunc, msg := animal(tarantula)
	if msg != "" {
		fmt.Println(msg)
	} else {
		total += tarFunc(2) 
	}

	fmt.Printf("el total en kg es %.2f", total) 
}
