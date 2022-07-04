package main

import (
	"errors"
	"fmt"
)

const (
	tarantula = "tarantula"
	hamster   = "hamster"
	gato      = "gato"
	perro     = "perro"
)

type Animal struct {
	tipo     string
	cantidad float64
}

func totalAlimento(animales ...Animal) (float64, error) {
	var total float64
	for _, animal := range animales {
		switch animal.tipo {
		case tarantula:
			total += totalTarantula(animal.cantidad)
		case hamster:
			total += totalHamster(animal.cantidad)
		case gato:
			total += totalGato(animal.cantidad)
		case perro:
			total += totalPerro(animal.cantidad)
		default:
			return 0, errors.New("El animal no se encuentra en la lista")
		}
	}

	return total, nil

}

func totalTarantula(cant float64) float64 {
	return 150 * cant
}

func totalHamster(cant float64) float64 {
	return 250 * cant
}

func totalGato(cant float64) float64 {
	return 5000 * cant
}

func totalPerro(cant float64) float64 {
	return 10000 * cant
}

func main() {
	cat := Animal{gato, 10}
	dog := Animal{perro, 20}
	spider := Animal{tarantula, 16}
	rat := Animal{hamster, 18}

	total, err := totalAlimento(cat, dog, spider, rat)
	if err != nil {
		fmt.Println("Ha ocurrido un error", err)
	}

	fmt.Println(total)
}
