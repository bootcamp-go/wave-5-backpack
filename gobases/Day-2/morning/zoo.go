package main

import "fmt"

const (
	dog       = "Dog"
	cat       = "Cat"
	tarantula = "Tarantula"
	hamster   = "Hamster"
)

func animalDog(amount float64) float64 {
	return 10000 * amount
}
func animalCat(amount float64) float64 {
	return 5000 * amount
}
func animalTarantula(amount float64) float64 {
	return 250 * amount
}
func animalHamster(amount float64) float64 {
	return 150 * amount
}
func Animal(name string) func(amount float64) float64 {
	switch name {
	case dog:
		return animalDog
	case cat:
		return animalCat
	case tarantula:
		return animalTarantula
	case hamster:
		return animalHamster
	default:
		return nil
	}
}
func main() {
	var amount float64
	dog1 := Animal(dog)
	cat1 := Animal(cat)
	amount += dog1(2)
	amount += cat1(1)
	fmt.Println(amount)
}
