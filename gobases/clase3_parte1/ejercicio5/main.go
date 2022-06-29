package main

import "fmt"

const (
	PERRO     = "perro"
	GATO      = "gato"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

func calcularAlimentos(cantidad, gramos int) int {
	return cantidad * gramos
}

func alimentosPerro(cantidad int) int {
	return calcularAlimentos(cantidad, 10000)
}

func alimentosGato(cantidad int) int {
	return calcularAlimentos(cantidad, 5000)
}

func alimentosHamster(cantidad int) int {
	return calcularAlimentos(cantidad, 250)
}

func alimentosTarantula(cantidad int) int {
	return calcularAlimentos(cantidad, 150)
}

func Animal(animal string) (func(cantidad int) int, string) {
	switch animal {
	case PERRO:
		return alimentosPerro, ""
	case GATO:
		return alimentosGato, ""
	case HAMSTER:
		return alimentosHamster, ""
	case TARANTULA:
		return alimentosTarantula, ""
	default:
		return nil, "Animal no valido"
	}
}

func main() {
	animalPerro, _ := Animal(PERRO)
	animalGato, _ := Animal(GATO)
	animalHamster, _ := Animal(HAMSTER)
	animalTarantula, _ := Animal(TARANTULA)

	var amount int
	amount += animalPerro(1)
	amount += animalGato(2)
	amount += animalHamster(4)
	amount += animalTarantula(1)

	fmt.Printf("Hay que comprar %d gramos de comida\n", amount)

	_, msg := Animal("Nada")
	fmt.Println(msg)
}
