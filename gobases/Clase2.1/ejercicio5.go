package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantura"
)

func dogFunc(quantity int) float64 {
	return float64(quantity) * 10
}

func catFunc(quantity int) float64 {
	return float64(quantity) * 5
}

func hamsterFunc(quantity int) float64 {
	return float64(quantity) * 0.250
}

func tarantulaFunc(quantity int) float64 {
	return float64(quantity) * 0.150
}

func Animal(option string) (func(quantity int) float64, error) {
	switch option {
	case dog:
		return dogFunc, nil
	case cat:
		return catFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	}
	return nil, errors.New("Animal " + option + " don't exist")
}

func main() {
	animalDog, msg1 := Animal(dog)
	animalCat, msg2 := Animal(cat)
	animalHams, msg3 := Animal(hamster)
	animalTaran, msg4 := Animal(tarantula)
	animalBird, msg5 := Animal("Bird")

	var amount float64 = 0
	if msg1 == nil {
		amount += animalDog(5)
	} else {
		fmt.Println(msg1)
	}

	if msg2 == nil {
		amount += animalCat(4)
	} else {
		fmt.Println(msg2)
	}

	if msg3 == nil {
		amount += animalHams(4)
	} else {
		fmt.Println(msg3)
	}

	if msg4 == nil {
		amount += animalTaran(1)
	} else {
		fmt.Println(msg4)
	}

	if msg5 == nil {
		amount += animalBird(5)
	} else {
		fmt.Println(msg5)
	}

	fmt.Printf("Total de comida: %v\n", amount)
}
