package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalPerro, msg := Animal(perro)
	animalGato, msg := Animal(gato)
	animalHamster, msg := Animal(hamster)
	animalTarantula, msg := Animal(tarantula)

	cantidadPerro := animalPerro(5)
	cantidadGato := animalGato(5)
	cantidadHamster := animalHamster(5)
	cantidadTarantula := animalTarantula(5)

	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Printf("Perros: %d kg\n", cantidadPerro)
		fmt.Printf("Gatos: %d kg\n", cantidadGato)
		fmt.Printf("Hamster: %d gramos\n", cantidadHamster)
		fmt.Printf("Tarantula: %d gramos\n", cantidadTarantula)

		fmt.Printf("En total se debe comprar: %d kg de comida\n", cantidadPerro+cantidadGato+cantidadHamster/1000+cantidadTarantula/1000)
	}
}

func Animal(animal string) (func(cantidad int) int, error) {
	switch animal {
	case perro:
		return perroFunc, nil
	case gato:
		return gatoFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		return func(cantidad int) int { return 0 }, errors.New("No existe el animal epecificado")
	}
}

func perroFunc(cantidad int) int {
	return cantidad * 10
}

func gatoFunc(cantidad int) int {
	return cantidad * 5
}

func hamsterFunc(cantidad int) int {
	return cantidad * 250
}

func tarantulaFunc(cantidad int) int {
	return cantidad * 150
}
