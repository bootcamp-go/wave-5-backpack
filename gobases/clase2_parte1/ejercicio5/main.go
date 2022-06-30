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

func perroFunc(valor int) float64 {
	a := float64(valor)

	return a * 10

}
func gatoFunc(valor int) float64 {
	a := float64(valor)

	return a * 5

}
func hamsterFunc(valor int) float64 {
	a := float64(valor)

	return a * 0.25

}
func tarantulaFunc(valor int) float64 {
	a := float64(valor)

	return a * 0.15

}

func animal(valor string) (func(valor2 int) float64, error) {
	switch valor {
	case "perro":
		return perroFunc, nil
	case "gato":
		return gatoFunc, nil
	case "hamster":
		return hamsterFunc, nil
	case "tarantula":
		return tarantulaFunc, nil

	}

	return nil, errors.New("animal no existe")

}
func main() {
	animalNombre := "hamster"
	res, err := animal(animalNombre)

	if err == nil {

		valor := 30
		cant := res(valor)
		fmt.Println("cantidad de alimento para", valor, animalNombre, "es de", cant, "Kg")

	} else {
		fmt.Println(err)
	}

}
