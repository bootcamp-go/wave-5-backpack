package main

import (
	"errors"
	"fmt"
)

const (
	DOG       = "dog"
	CAT       = "cat"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func main() {
	dog, msg := Animal(DOG)
	cat, msg := Animal(CAT)
	hamster, msg := Animal(HAMSTER)
	tarantula, msg := Animal(TARANTULA)

	if msg != nil {
		fmt.Println("Hubo un error", msg)
	}

	fmt.Printf("Cantidad de alimento para 2 perros: %v kg \n", dog(2))
	fmt.Printf("Cantidad de alimento para 3 gatos: %v kg \n", cat(3))
	fmt.Printf("Cantidad de alimento para 4 hamster: %v kg \n", hamster(4))
	fmt.Printf("Cantidad de alimento para 5 tarantulas: %v kg \n", tarantula(5))

}

func animalDog(cant int) float64 {
	return float64(cant) * 10
}

func animalCat(cant int) float64 {
	return float64(cant) * 5
}

func animalHamster(cant int) float64 {
	return float64(cant) * (250.0 / 100.0)
}

func animalTarantula(cant int) float64 {
	return float64(cant) * (150.0 / 100.0)
}

func Animal(animal string) (func(cantidad int) float64, error) {
	switch animal {
	case TARANTULA:
		return animalTarantula, nil
	case HAMSTER:
		return animalHamster, nil
	case DOG:
		return animalDog, nil
	case CAT:
		return animalCat, nil
	default:
		return nil, errors.New("No existe el animal: " + animal)
	}
}
