package main

import (
	"errors"
	"fmt"
)

func amountTarantula(numAnimals float64) float64 {
	foodKg := 150.0 / 1000
	return foodKg * numAnimals
}
func amountHamster(numAnimals float64) float64 {
	foodKg := 250.0 / 1000
	return foodKg * numAnimals
}
func amountDog(numAnimals float64) float64 {
	return 10 * numAnimals
}

func amountCat(numAnimals float64) float64 {
	return 5 * numAnimals
}

func orquestador(numAnimals float64, amountFoodOneAnimal func(numAnimals float64) float64) float64 {

	return amountFoodOneAnimal(numAnimals)

}

func amountFood(numAnimals float64, typeAnimal string) (float64, error) {
	switch typeAnimal {
	case "tarantula":
		return orquestador(numAnimals, amountTarantula), nil
	case "hamster":
		return orquestador(numAnimals, amountHamster), nil
	case "perro":
		return orquestador(numAnimals, amountDog), nil
	case "gato":
		return orquestador(numAnimals, amountCat), nil

	}
	return 0, errors.New("El animal no se encuentra")
}

func main() {

	res, err := amountFood(3, "aguila")
	if err != nil {
		fmt.Printf("Ocurrio un error: %v \n", err)
	} else {
		fmt.Printf("El resulto de la operación es: %v \n", res)
	}
}
