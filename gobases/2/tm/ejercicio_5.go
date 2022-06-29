package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	var amount float64

	dogs := 1
	cats := 1
	hamsters := 1
	tarantulas := 1

	animalDog, msg := animal(dog)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalDog(dogs)
	}

	animalCat, msg := animal(cat)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalCat(cats)
	}

	animalHamster, msg := animal(hamster)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalHamster(hamsters)
	}

	animalTarantula, msg := animal(tarantula)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalTarantula(tarantulas)
	}

	fmt.Printf("Alimento necesario: %.3f Kg\n", amount)
}

func animal(a string) (func(quantity int) (amount float64), error) {
	switch a {
	default:
		return nil, errors.New("Animal no disponible por el momento: " + a)
	case "dog":
		return func(quantity int) (amount float64) {
			amount = float64(quantity) * 10
			return amount
		}, nil

	case "cat":
		return func(quantity int) (amount float64) {
			amount = float64(quantity) * 5
			return amount
		}, nil

	case "hamster":
		return func(quantity int) (amount float64) {
			amount = float64(quantity) * 0.250
			return amount
		}, nil

	case "tarantula":
		return func(quantity int) (amount float64) {
			amount = float64(quantity) * 0.150
			return amount
		}, nil
	}
}
