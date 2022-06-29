package main

import (
	"errors"
	"fmt"
)

const (
	perro            string  = "Perro"
	comida_perro     float64 = 10
	gato             string  = "Gato"
	comida_gato      float64 = 5
	hamster          string  = "Hamnster"
	comida_hamster   float64 = 0.25
	tarantula        string  = "Tarantula"
	comida_tarantula float64 = 0.15
	other            string  = "Ratón"
)

func main() {
	amountTotal := calcAmount()
	fmt.Printf("En total se necesitan %f Kg de alimento para todos los animales\n", amountTotal)
}

func Animal(animal string) (func(cant int) float64, error) {
	switch animal {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	}
	return nil, errors.New("El animal " + animal + " no existe en el refugio")
}

func animalPerro(cant int) float64 {
	var amount float64 = float64(cant) * comida_perro
	return amount
}

func animalGato(cant int) float64 {
	var amount float64 = float64(cant) * comida_gato
	return amount
}

func animalHamster(cant int) float64 {
	var amount float64 = float64(cant) * comida_hamster
	return amount
}

func animalTarantula(cant int) float64 {
	var amount float64 = float64(cant) * comida_tarantula
	return amount
}

func calcAmount() float64 {
	var amountTotal float64 = 0
	var amount float64
	perroFunc, msg := Animal(perro)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount = perroFunc(2)
		amountTotal += amount
		fmt.Printf("Los perros necesitan %f Kg de alimento\n", amount)
	}

	gatoFunc, msg := Animal(gato)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount = gatoFunc(3)
		amountTotal += amount
		fmt.Printf("Los gatos necesitan %f Kg de alimento\n", amount)
	}

	hamsterFunc, msg := Animal(hamster)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount = hamsterFunc(10)
		amountTotal += amount
		fmt.Printf("Los hamster necesitan %f Kg de alimento\n", amount)
	}

	tarantulaFunc, msg := Animal(tarantula)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount = tarantulaFunc(20)
		amountTotal += amount
		fmt.Printf("Las tarantulas necesitan %f Kg de alimento\n", amount)
	}

	otherFunc, msg := Animal(other)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount = otherFunc(5)
		amountTotal += amount
		fmt.Printf("Los otros animales necesitan %f Kg de alimento\n", amount)
	}

	return amountTotal
}
