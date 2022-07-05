package main

import (
	"errors"
	"fmt"
)

const (
	PERRO   = "perro"
	GATO    = "gato"
	HAMSTER = "hamster"
	ARANA   = "arana"
)

var (
	cantPerro   = 1
	cantGato    = 2
	cantHamster = 1
	cantArana   = 1
)

func main() {

	animalPerro, errPerro := animal(PERRO)
	if errPerro != nil {
		fmt.Println(errPerro)
	} else {
		fmt.Printf("La cantidad de comida para %d %s es: %.2fkg.\n", cantPerro, PERRO, animalPerro(cantPerro))
	}

	animalGato, errGato := animal(GATO)
	if errGato != nil {
		fmt.Println(errGato)
	} else {
		fmt.Printf("La cantidad de comida para %d %s es: %.2fkg.\n", cantGato, GATO, animalGato(cantGato))
	}

	animalHamster, errHamster := animal(HAMSTER)
	if errHamster != nil {
		fmt.Println(errHamster)
	} else {
		fmt.Printf("La cantidad de comida para %d %s es: %.2fkg\n", cantHamster, HAMSTER, animalHamster(cantHamster))
	}

	animalArana, errArana := animal(ARANA)
	if errPerro != nil {
		fmt.Println(errArana)
	} else {
		fmt.Printf("La cantidad de comida para %d %s es: %.2fkg\n", cantArana, ARANA, animalArana(cantArana))
	}

	_, errAnimalInvalido := animal("")
	if errAnimalInvalido != nil {
		fmt.Println(errAnimalInvalido)
	}

}

func animal(animal string) (func(q int) float64, error) {
	switch animal {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case HAMSTER:
		return animalHamster, nil
	case ARANA:
		return animalArana, nil
	default:
		return nil, errors.New("Ha ingresado un nombre no válido")
	}
}

func animalPerro(cantidad int) float64 {
	return 10 * float64(cantidad)
}

func animalGato(cantidad int) float64 {
	return 5 * float64(cantidad)
}

func animalHamster(cantidad int) float64 {
	return 0.25 * float64(cantidad)
}

func animalArana(cantidad int) float64 {
	return 0.15 * float64(cantidad)
}
