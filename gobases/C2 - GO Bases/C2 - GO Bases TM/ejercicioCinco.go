package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("Amount of food: %f needed\n", quantityFood())
}

const(
	//Animalitos
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"

	//Comida
	DogFood = 10
	CatFood = 5
	HamsterFood = 250
	TarantulaFood = 150
)

//finalmente asigno los valores y el amount para mostrar la totalidad de comida
//requerida
func quantityFood() float64 {
	var amount float64

	animalDog, msg := Animal(dog)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalDog(2)
	}

	animalCat, msg := Animal(cat)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalCat(3)
	}

	animalHamster, msg := Animal(hamster)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalHamster(10)
	}

	animalTarantula, msg := Animal(tarantula)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalTarantula(20)
	}

	return amount
}

//Funciones por cada animalito
func animalDog(cantidad int) float64 {
	return calculate(cantidad, DogFood, "dogs","kg")
}

func animalCat(cantidad int) float64 {
	return calculate(cantidad, CatFood, "cats","kg")
}

func animalHamster(cantidad int) float64 {
	return calculate(cantidad, HamsterFood, "hamsters","gram")
}

func animalTarantula(cantidad int) float64 {
	return calculate(cantidad, TarantulaFood, "tarantulas","gram")
}

//Cálculo mediante abstracción de las cantidades requeridas por cada uno
func calculate(cantidad int, animalFood int, typeAnimal string, cantFood string) float64{
	total := float64(cantidad) * float64(animalFood)
	fmt.Println(typeAnimal, "needs", total,cantFood,"of food")
	return total
}

//switch para elegir el tipo de animalito
func Animal(animal string) (func(cant int) float64, error) {
	switch animal {
	case dog:
		return animalDog, nil
	case cat:
		return animalCat, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	}
	return nil, errors.New("There isn't that kind of animals in the shelter")
}