package main

import (
	"errors"
	"fmt"
)

const (
	PERRO     = "perro"
	GATO      = "gato"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

type animalValido func(int) float64

func obtenerComidaPerro(numero int) float64 {
	return float64(numero) * 10
}

func obtenerComidaGato(numero int) float64 {
	return float64(numero * 5)
}

func obtenerComidaHamster(numero int) float64 {
	return float64(numero) * 0.25
}

func obtenerComidaTarantula(numero int) float64 {
	return float64(numero) * 0.15
}

func Animal(operacion string) (animalValido, error) {
	switch operacion {
	default:
		return nil, errors.New("Animal no valido")
	case PERRO:
		return obtenerComidaPerro, nil
	case GATO:
		return obtenerComidaGato, nil
	case HAMSTER:
		return obtenerComidaHamster, nil
	case TARANTULA:
		return obtenerComidaTarantula, nil
	}
}

func calculadora(operacion string, numero int) (float64, error) {
	operacionSeleccionada, err := Animal(operacion)
	if err != nil {
		return 0, err
	}
	return operacionSeleccionada(numero), nil
}

func main() {
	operacion := "tarantula"
	numero := 1
	res, err := calculadora(operacion, numero)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
