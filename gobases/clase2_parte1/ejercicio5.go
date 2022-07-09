package main

import "fmt"

const(
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	spider = "spider"
)

func main()  {

	var cantidad float64

	alimentodog, err := animal(dog)
	if err != "" {
		fmt.Println(err)
	} else {
		cantidad += alimentodog(5) // 50kg
	}

	alimentocat, err := animal(cat)
	if err != "" {
		fmt.Println(err)
	} else {
		cantidad += alimentocat(5) // 10kg
	}

	alimentohamster, err := animal(hamster)
	if err != "" {
		fmt.Println(err)
	} else {
		cantidad += alimentohamster(4) // 0,5kg
	}

	alimentospider, err := animal(spider)
	if err != "" {
		fmt.Println(err)
	} else {
		cantidad += alimentospider(2) // 0,3kg
	}

	fmt.Println("Alimento total para los animales", cantidad, "Kg")
}

func animal(animal string) (func(int) float64, string) {
	switch animal {
	case dog:
		return AlimentoDog, ""
	case cat:
		return AlimentoCat, ""
	case hamster:
		return AlimentoHamster, ""
	case spider:
		return AlimentoSpider, ""
	default:
		return nil, "El animal ingresado no existe"
	}
}

func AlimentoDog(alimento int) float64 {
	return 10 * float64(alimento)
}
func AlimentoCat(alimento int) float64 {
	return 5 * float64(alimento)
}
func AlimentoHamster(alimento int) float64 {
	return (250 * float64(alimento)) / 1000
}
func AlimentoSpider(alimento int) float64 {
	return (150 * float64(alimento)) / 1000
}
