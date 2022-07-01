package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	TARANTULA = "tarantula"
	HAMSTER   = "hamster"
)

const (
	COMIDA_PERROS     = 10
	COMIDA_GATOS      = 5
	COMIDA_TARANTULAS = 0.15
	COMIDA_HAMSTER    = 0.25
)

func main() {
	animalPerro, msg := Animal(PERRO)
	animalGato, msg := Animal(GATO)
	animalTarantula, msg := Animal(TARANTULA)
	animalHamster, msg := Animal(HAMSTER)
	//_, msg = Animal("INVALIDA")

	var amount float32

	if msg != nil {
		fmt.Println(msg)
		return
	}

	amount += animalPerro(5)      //50kg
	amount += animalGato(1)       //5kg
	amount += animalTarantula(20) //3kg
	amount += animalHamster(4)    //1kg

	fmt.Println("Cantidad total de alimento: ", amount)
}

func Animal(animal string) (func(int) float32, error) {
	switch animal {
	case PERRO:
		return comidaPerros, nil
	case GATO:
		return comidaGatos, nil
	case TARANTULA:
		return comidaTarantulas, nil
	case HAMSTER:
		return comidaHamster, nil
	}

	return nil, errors.New("El animal ingresado no existe")

}

func comidaPerros(cantidadPerros int) float32 {
	return float32(cantidadPerros) * COMIDA_PERROS
}

func comidaGatos(cantidadGatos int) float32 {
	return float32(cantidadGatos) * COMIDA_GATOS
}

func comidaTarantulas(cantidadTarantulas int) float32 {
	return float32(cantidadTarantulas) * COMIDA_TARANTULAS
}

func comidaHamster(cantidadHamsters int) float32 {
	return float32(cantidadHamsters) * COMIDA_HAMSTER
}
