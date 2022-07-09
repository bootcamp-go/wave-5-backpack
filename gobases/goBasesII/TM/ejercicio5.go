package main

import (
	"errors"
	"fmt"
)

const (
	Perro     string = "perro"
	Gato      string = "gato"
	Hamster   string = "hamster"
	Tarantula string = "tarantula"
)

func Animal(animal string) (calAlimento func(cantidad int) float64, err error) {
	switch animal {
	case Perro:
		return calPerro, nil
	case Gato:

		return calGato, nil
	case Hamster:
		return calHamster, nil
	case Tarantula:
		return calTarantula, nil
	}
	return nil, errors.New("No hay una operacion definida para resolver la solicitud")
}

func calPerro(cant int) float64 {
	return float64(cant) * 10
}
func calGato(cant int) float64 {
	return float64(cant) * 5
}
func calHamster(cant int) float64 {
	return float64(cant) * 0.25
}
func calTarantula(cant int) float64 {
	return float64(cant) * 0.15
}
func main() {
	var amount float64
	calcularPerro, msg := Animal(Perro)
	if msg != nil {
		fmt.Println(msg)
	} else {
		cPerro := calcularPerro(10)
		amount += cPerro
		fmt.Printf("Perro comida: %.2fKg // cantidad: %d\n", cPerro, 10)
	}
	calcularGato, msg := Animal(Gato)
	if msg != nil {
		fmt.Println(msg)
	} else {
		cGato := calcularGato(10)
		amount += cGato
		fmt.Printf("Perro comida: %.2fKg // cantidad: %d\n", cGato, 10)
	}
	calcularHamster, msg := Animal(Hamster)
	if msg != nil {
		fmt.Println(msg)
	} else {
		cHamster := calcularHamster(10)
		amount += cHamster
		fmt.Printf("Perro comida: %.2fKg // cantidad: %d\n", cHamster, 10)
	}
	calcularTara, msg := Animal(Tarantula)
	if msg != nil {
		fmt.Println(msg)
	} else {
		cTara := calcularTara(10)
		amount += cTara
		fmt.Printf("Perro comida: %.2fKg // cantidad: %d\n", cTara, 10)
	}
	fmt.Println(amount)
}
