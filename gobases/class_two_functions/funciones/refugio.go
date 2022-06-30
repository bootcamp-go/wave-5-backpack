package main

import "fmt"

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func main() {
	animalDog := Animal(dog)
	dog, msg := animalDog(5)
	fmt.Println("Necesitamos ", dog, "kg.", msg)

	animalCat := Animal(cat)
	cat, msgCat := animalCat(15)
	fmt.Println("Necesitamos ", cat, "kg.", msgCat)
}

func Animal(animal string) func(amount float32) (float32, string) {
	switch animal {
	case dog:
		return animalDog
	case cat:
		return animalCat
	case spider:
		return animalSpider
	case hamster:
		return animalHamster
	default:
		return noFoundAnimal
	}
}

func noFoundAnimal(amount float32) (total float32, msg string) {
	total = 0
	msg = "No está registrado ese animalito :("
	return
}
func animalDog(amount float32) (total float32, msg string) {
	msg = "¡Estos son los kg que necesitamos en total por los perritos!"
	total = (amount) * (10)
	return
}

func animalCat(amount float32) (total float32, msg string) {
	msg = "¡Estos son los kg que necesitamos en total por los gatitos!"
	total = (amount) * (5)
	return
}

func animalSpider(amount float32) (total float32, msg string) {
	msg = "¡Estos son los kg que necesitamos en total por las arañitas!"
	total = (amount) * (.250)
	return
}

func animalHamster(amount float32) (total float32, msg string) {
	msg = "¡Estos son los kg que necesitamos en total por los hamsters!"
	total = (amount) * (.150)
	return
}
