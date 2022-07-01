package main

import (
	"fmt"
)

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
// Por el momento solo tienen tarántulas, hamsters, perros, y gatos,
// pero se espera que puedan haber muchos más animales que refugiar.

// perro necesitan 10 kg de alimento
// gato 5 kg
// Hamster 250 gramos.
// Tarántula 150 gramos.

// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con
// el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo
// de animal especificado.

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	otro    = "otro"
)

func animalDog(cantidad float64) float64 {

	return cantidad * 10
}

func animalHamster(cantidad float64) float64 {

	return cantidad * 0.250
}

func animalTarantula(cantidad float64) float64 {

	return cantidad * 0.150
}

func animalCat(cantidad float64) float64 {

	return cantidad * 5
}

func Animal(animal string) (func(cantidad float64) float64, string) {
	switch animal {
	case "dog":
		return animalDog, "Comida para perro"
	case "cat":
		return animalCat, "Comida para gato"
	default:
		panic("No tenemos a ese animal definido")

	}

}

func main() {

	var amount float64

	amount += animalDog(5)
	animalDog, msg := Animal(dog)
	fmt.Println(animalDog(amount), msg)

	animalCat, msg := Animal(cat)
	amount += animalCat(8)
	fmt.Println(animalCat(amount), msg)

	animalNo, msg := Animal(otro)
	amount += animalNo(8)
	fmt.Println(animalNo(amount), msg)
}
