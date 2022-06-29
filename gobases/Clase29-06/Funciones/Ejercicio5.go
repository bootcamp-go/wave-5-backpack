package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "PERRO"
	GATO      = "GATO"
	TARANTULA = "TARANTULA"
	HAMSTER   = "HAMSTER"
)

func animalPerro(cantidad int) int {
	return (cantidad * 10)
}

func animalGato(cantidad int) int {
	return (cantidad * 5)
}

func animalTarantula(cantidad int) int {
	return (cantidad * 250 / 1000)
}

func animalHamster(cantidad int) int {
	return (cantidad * 150 / 1000)
}

func animal(tipo string) (func(values int) int, error) {
	switch tipo {
	case PERRO:
		return animalPerro, nil
	case GATO:
		return animalGato, nil
	case TARANTULA:
		return animalTarantula, nil
	case HAMSTER:
		return animalHamster, nil
	}
	return nil, errors.New("EL ANIMAL NO EXISTE")
}

func main() {

	r, _ := animal(PERRO)
	j := r(4)
	fmt.Printf("Se necesitan %v KGs", j)

}
