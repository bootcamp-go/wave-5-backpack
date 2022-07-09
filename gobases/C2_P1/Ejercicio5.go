package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "perro"
	cat       = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animal(animal string, value int) (int, error) {
	var resultFood int
	switch animal {
	case dog:
		resultFood = int(value * 10)
		return fmt.Printf("La cantidad de comida necesaria para %v %ss es: %v kg\n", value, animal, resultFood)
	case cat:
		resultFood = int(value * 5)
		return fmt.Printf("La cantidad de comida necesaria para %v %ss es: %v kg\n", value, animal, resultFood)
	case hamster:
		resultFood = int(250 * float64(value))
		return fmt.Printf("La cantidad de comida necesaria para %v %ss es: %v g\n", value, animal, resultFood)
	case tarantula:
		resultFood = int(150 * float64(value))
		return fmt.Printf("La cantidad de comida necesaria para %v %ss es: %v g\n", value, animal, resultFood)
	default:
		return 0, errors.New("El animal ingresado no existe en la lista")
	}
}

func main() {
	animal, err := animal(hamster, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(animal)
	}
}
