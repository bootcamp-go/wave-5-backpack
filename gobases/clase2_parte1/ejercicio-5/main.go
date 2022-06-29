package main

import (
	"errors"
	"fmt"
)

const (
	TARANTULA = "Tarántula"
	HAMSTER   = "Hamster"
	PERRO     = "Perro"
	GATO      = "Gato"
)

func main() {

	funcAnimal, msg := Animal(TARANTULA)
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Cantidad de alimento para 5 tarantulas:", funcAnimal(5), "kg")
	}

	funcAnimal, msg = Animal(HAMSTER)
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Cantidad de alimento para 5 hamsters:", funcAnimal(5), "kg")
	}

	funcAnimal, msg = Animal(PERRO)
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Cantidad de alimento para 5 perros:", funcAnimal(5), "kg")
	}

	funcAnimal, msg = Animal(GATO)
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Cantidad de alimento para 5 gatos:", funcAnimal(5), "kg")
	}

	funcAnimal, msg = Animal("Conejo")
	if msg != nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Cantidad de alimento para 5 conejos:", funcAnimal(5), "kg")
	}

}

func Animal(animal string) (func(cantidad int) float64, error) {
	switch animal {
	case TARANTULA:
		return animalTarantula, nil
	case HAMSTER:
		return animalHamster, nil
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	default:
		return nil, errors.New("No existe el animal: " + animal)
	}
}

func animalTarantula(cantidad int) float64 {
	// Resultado dado en kg
	return float64(cantidad) * (150.0 / 1000.0)
}

func animalHamster(cantidad int) float64 {
	// Resultado dado en kg
	return float64(cantidad) * (250.0 / 1000.0)
}

func animalPerro(cantidad int) float64 {
	// Resultado dado en kg
	return float64(cantidad) * 5
}

func animalGato(cantidad int) float64 {
	// Resultado dado en kg
	return float64(cantidad) * 10
}
