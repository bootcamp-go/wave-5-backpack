package main

import "fmt"

const (
	CAT       = "CAT"
	DOG       = "DOG"
	HAMSTER   = "HAMSTER"
	TARANTULA = "TARANTULA"
)

func foodDog(quantity_animal int) float32 {
	return float32(quantity_animal * 10)
}

func foodCat(quantity_animal int) float32 {
	return float32(quantity_animal * 5)
}

func foodHamster(quantity_animal int) float32 {
	return float32(quantity_animal*250) / 1000
}

func foodTarantula(quantity_animal int) float32 {
	return float32(quantity_animal*150) / 1000
}

func foodAnimal(animal string) func(int) float32 {
	switch animal {
	case "DOG":
		return foodDog
	case "CAT":
		return foodCat
	case "HAMSTER":
		return foodHamster
	case "TARANTULA":
		return foodTarantula
	default:
		return nil
	}
}

func main() {
	food_cat := foodAnimal(CAT)
	food_dog := foodAnimal(DOG)
	food_ham := foodAnimal(HAMSTER)
	food_tar := foodAnimal(TARANTULA)

	fmt.Printf("Dog: %.2f\n", food_dog(3))
	fmt.Printf("Cat: %.2f\n", food_cat(5))
	fmt.Printf("Hamster: %.2f\n", food_ham(3))
	fmt.Printf("Tarantula: %.2f\n", food_tar(3))
}
