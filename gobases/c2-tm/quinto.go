package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func fdog() {
	fmt.Println("necesita 10 kg de alimento")
}

func fcat() {
	fmt.Println("necesita 5 kg de alimento")
}

func fhamster() {
	fmt.Println("necesita 250g de alimento")
}

func fspider() {
	fmt.Println("necesita 150g de alimento")
}

func Animal(typeAnimal string) (func(), error) {

	switch typeAnimal {
	case dog:
		return fdog, nil
	case cat:
		return fcat, nil
	case hamster:
		return fhamster, nil
	case spider:
		return fspider, nil
	}

	return nil, errors.New("El tipo de animal no existe")
}

func main() {
	animalDog, err := Animal(dog)
	if err != nil {
		fmt.Println("no se pudo encontrar el tipo de animal")
		return
	}

	animalCat, err := Animal(cat)
	if err != nil {
		fmt.Println("no se pudo encontrar el tipo de animal")
		return
	}

	animalDog()
	animalCat()
}
