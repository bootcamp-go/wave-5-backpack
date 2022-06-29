package main

import (
	"errors"
	"fmt"
)

func Animal(animal string) (func(cantidad int) float64, error) {
	comida := map[string]float64{
		"perro":     10000,
		"gato":      5000,
		"hamster":   250,
		"tarantula": 150,
	}

	gramos, exists := comida[animal]

	if exists {
		return func(cantidad int) float64 {
			return float64(cantidad) * gramos
		}, nil
	}

	return nil, errors.New("Animal no valido")
}

func main() {
	animalDog, err := Animal("perro")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	animalCat, err := Animal("gato")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	var amount float64
	amount += animalDog(5)
	amount += animalCat(8)

	fmt.Printf("Cantidad de comida: %f gramos\n", amount)

}
