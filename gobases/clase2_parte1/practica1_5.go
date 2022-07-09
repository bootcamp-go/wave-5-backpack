package main

import (
	"fmt"
	"errors"
)

//Ejercicio 5 - Calcular cantidad de alimento
const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
    tarantula = "tarantula"
)

func cantComidaPerro(cantidad int)  int {
	alimento := 10
	return cantidad * alimento
}

func cantComidaGato(cantidad int) int {
	alimento := 5
	return cantidad * alimento
}

func cantComidaHamster(cantidad int) int {
	alimento := 250
	return cantidad * alimento
}

func cantComidaTarantula(cantidad int) int {
	alimento := 150
	return cantidad * alimento
}

func Animal(tipoAnimal string) (func(int) int, error)  {
	switch tipoAnimal {
	case perro:
		return cantComidaPerro, nil
	case gato:
		return cantComidaGato, nil
	case hamster:
		return cantComidaHamster, nil
	case tarantula:
		return cantComidaTarantula, nil
	default:
		return nil, errors.New("Ese animal no existe")
	}
}

func main()  {
	
	animalPerro, err := Animal(perro)
	if err != nil {
		fmt.Println(err)
		return
	}

	animalGato, err := Animal(gato)
	if err != nil {
		fmt.Println(err)
		return
	}

	animalHamster, err := Animal(hamster)
	if err != nil {
		fmt.Println(err)
		return
	}

	animalTarantula, err := Animal(tarantula)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("La cantidad de alimento que se necesita para alimentar a los perros es de", animalPerro(3), "kg.")
	fmt.Println("La cantidad de alimento que se necesita para alimentar a los gatos es de", animalGato(6), "kg.")
	fmt.Println("La cantidad de alimento que se necesita para alimentar a los hamsters es de", animalHamster(8), "gramos.")
	fmt.Println("La cantidad de alimento que se necesita para alimentar a las tarantulas es de", animalTarantula(2), "gramos.")

}