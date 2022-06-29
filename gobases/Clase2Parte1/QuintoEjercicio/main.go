package main

import "fmt"

const (
	Dog     = "dog"
	Cat     = "cat"
	Spider  = "tarantula"
	Hamster = "hamster"
)

func animalDog(dog int) float64 {
	return float64(dog) * 10
}
func animalCat(cat int) float64 {
	return float64(cat) * 5
}
func animalSpider(spider int) float64 {
	return float64(spider) * 0.150
}
func animalHamster(hamster int) float64 {
	return float64(hamster) * 0.250
}

func Animal(animal string) (func(num int) float64, string) {
	switch animal {
	case Dog:
		return animalDog, ""
	case Cat:
		return animalCat, ""
	case Spider:
		return animalSpider, ""
	case Hamster:
		return animalHamster, ""
	default:
		return nil, "El animal no existe"
	}
}

func main() {
	animalDog, msg := Animal(Dog)
	animalCat, msg2 := Animal(Cat)
	var amount float64
	amount += animalDog(5)
	amount += animalCat(8)
	if msg != "" && msg2 != "" {
		fmt.Printf("Animal 1: %v , Animal 2: %v \n", msg, msg2)
	} else {
		fmt.Printf("La cantidad de comida necesaria es: %v \n", amount)
	}
}
