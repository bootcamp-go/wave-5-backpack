package main

import (
	"fmt"
	"errors"
)

const(
	DOG = "dog"
	CAT = "cat"
	HAMSTER = "hamster"
	TARANTULA = "tarantula"
)

func cantPerro(cant int) float32 {
	return float32(cant) * 10;
}

func cantGato(cant int) float32 {
	return float32(cant) * 5
}

func cantHamster(cant int) float32 {
	return float32(cant) * 0.25
}

func cantTarantula(cant int) float32 {
	return float32(cant) * 0.15
}

func Animal(animal string) (func(cantidad int) float32, error) {
	switch animal {
		case DOG:
			return cantPerro, nil
		case CAT:
			return cantGato, nil
		case HAMSTER:
			return cantHamster, nil
		case TARANTULA:
			return cantTarantula, nil
		default:
			return nil, errors.New("No tenemos este animal")
	}
} 

func main() {
	dogFunc, err := Animal(DOG)
	catFunc, err := Animal(CAT)
	hamsterFunc, err := Animal(HAMSTER)
	tarantulaFunc, err := Animal(TARANTULA)
	if err == nil {
		dogValue := dogFunc(6)
		catValue := catFunc(8)
		hamsterValue := hamsterFunc(9)
		tarantulaValue := tarantulaFunc(9)
		fmt.Println(dogValue, catValue, hamsterValue, tarantulaValue)
	} else {
		fmt.Println(err)
	}
}