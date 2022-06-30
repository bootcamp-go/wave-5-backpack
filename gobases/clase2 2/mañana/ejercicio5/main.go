package main

import "fmt"

/*Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
haber muchos más animales que refugiar.

	1. perro necesitan 10 kg de alimento
	2. gato 5 kg
	3. Hamster 250 gramos.
	4. Tarántula 150 gramos.

Se solicita:

	1. Implementar una función Animal que reciba como parámetro un valor de tipo texto con el
	animal especificado y que retorne una función y un mensaje (en caso que no exista el animal).
	2. Una función para cada animal que calcule la cantidad de alimento en base a la cantidad
	del tipo de animal especificado.*/

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarántula"
)

func animalDog(dogs int) float64 {
	return 10 * float64(dogs)
}

func animalCat(cats int) float64 {
	return 5 * float64(cats)
}

func animalTarantula(tarantula int) float64 {
	return (150 * float64(tarantula)) / 1000 // gramos
}

func animalHamster(hamster int) float64 {
	return (250 * float64(hamster)) / 1000 // gramos
}

func Animal(animal string) (func(int) float64, string) {
	switch animal {
	case dog:
		return animalDog, ""
	case cat:
		return animalCat, ""
	case hamster:
		return animalHamster, ""
	case tarantula:
		return animalTarantula, ""
	default:
		return nil, "Este animal no se encuentra registrado."
	}
}

func main() {
	var cantidad float64

	animaldog, msg := Animal(dog)
	if msg != "" {
		fmt.Println(msg)
	} else {
		cantidad += animaldog(2) // 20 kgs
	}

	animalcat, msg := Animal(cat)
	if msg != "" {
		fmt.Println(msg)
	} else {
		cantidad += animalcat(5) // cantidad + 25 kgs
	}

	animalhamster, msg := Animal(hamster)
	if msg != "" {
		fmt.Println(msg)
	} else {
		cantidad += animalhamster(4) // cantidad + 1 kg
	}

	animaltarantula, msg := Animal(tarantula)
	if msg != "" {
		fmt.Println(msg)
	} else {
		cantidad += animaltarantula(2) // cantidad + 0,3 kg
	}

	fmt.Println("✅ Total de alimentos necesarios", cantidad, "Kg") // 46,3 Kg
}
