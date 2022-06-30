package main

import (
	"fmt"
	"errors"
)

func main(){
animalDog := Animal(perro)
fmt.Println(animalDog, )
}

const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
)

func Animal(animal string)int{
	switch animal {
	case perro: 
		return perroFunc()
	case gato:
		return gatoFunc()
	case hamster: 
		return hamsterFunc()
	case tarantula: 
		return tarantulaFunc()
	}
	errors.New("No se encontro el animal especificado")
	return 0
}

func perroFunc()int{
	return 10
}
func gatoFunc()int{
	return 5
}
func hamsterFunc()int{
	return 250
}
func tarantulaFunc()int{
	return 150
}