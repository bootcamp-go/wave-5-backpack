package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = 10000
	GATO      = 5000
	HAMSTER   = 250
	TARANTULA = 150
)

func calcularCantidadDeAlimento(animal string) {
	//Ejercicio 5
	alimentoAnimal, err := cantidadAlimento(animal)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Cantidad alimento para %s: %.2f grs.", animal, alimentoAnimal(2))
	}
}

func cantidadAlimento(animal string) (func(cantidadAnimal int) float64, error) {
	//Ejercicio 5
	switch animal {
	case "perro":
		return alimentoPerro, nil
	case "gato":
		return alimentoGato, nil
	case "hamster":
		return alimentoHamster, nil
	case "tarantula":
		return alimentoTarantula, nil
	default:
		return nil, errors.New("El animal no existe")

	}
}

func alimentoPerro(cantidadPerros int) float64 {
	return float64(PERRO * cantidadPerros)
}

func alimentoGato(cantidadGatos int) float64 {
	return float64(GATO * cantidadGatos)
}

func alimentoHamster(cantidadHamster int) float64 {
	return float64(HAMSTER * cantidadHamster)
}

func alimentoTarantula(cantidadTarantulas int) float64 {
	return float64(TARANTULA * cantidadTarantulas)
}
