package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(animal string) (func(cantidad int) float32, error) {
	switch animal {
	case perro:
		return alimentoPerro, nil
	case gato:
		return alimentoGato, nil
	case hamster:
		return alimentoHamster, nil
	case tarantula:
		return alimentoTarantula, nil
	}
	return nil, errors.New("No existe el animal.")
}

func alimentoPerro(cantidad int) float32 {
	alimento := float32(cantidad) * 10
	return alimento
}

func alimentoGato(cantidad int) float32 {
	alimento := float32(cantidad) * 5
	return alimento
}

func alimentoHamster(cantidad int) float32 {
	alimento := float32(cantidad) * 0.25
	return alimento
}

func alimentoTarantula(cantidad int) float32 {
	alimento := float32(cantidad) * 0.15
	return alimento
}

func main() {
	res, err := Animal(perro)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("La cantidad de alimento para los perros es ", res(2), "Kg")
	}

	res2, err2 := Animal(gato)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("La cantidad de alimento para los gatos es ", res2(3), "Kg")
	}

	res3, err3 := Animal(hamster)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println("La cantidad de alimento para los hamster es ", res3(4), "Kg")
	}

	res4, err4 := Animal(tarantula)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println("La cantidad de alimento para los perros es ", res4(5), "Kg")
	}
}
