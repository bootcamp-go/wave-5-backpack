package main

import "fmt"

// Ejercicio 5 - Calcular cantidad de alimento

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
// Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
// haber muchos más animales que refugiar.

// 1. perro necesitan 10 kg de alimento
// 2. gato 5 kg
// 3. Hamster 250 gramos.
// 4. Tarántula 150 gramos.

// Se solicita:

// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado
// y que retorne una función y un mensaje (en caso que no exista el animal)
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo
// de animal especificado.

const (
	dog = "dog"
	cat = "cat"
	ham = "ham"
	tar = "tar"
)

func animalDog(cantidad int) float64 {
	return float64(10 * cantidad)
}

func animalCat(cantidad int) float64 {
	return float64(5 * cantidad)
}

func animalHam(cantidad int) float64 {
	return 0.25 * float64(cantidad)
}
func animalTar(cantidad int) float64 {
	return 0.15 * float64(cantidad)
}

func Animal(animal string) (func(cantidad int) float64, string) {
	switch animal {
	case dog:
		return animalDog, ""
	case cat:
		return animalCat, ""
	case ham:
		return animalHam, ""
	case tar:
		return animalTar, ""
	default:
		return nil, "No existe el animal"
	}
}

func main() {
	fmt.Println("Ejercicio 5 - Calcular cantidad de alimento")
	fmt.Println("")

	animalDog, msgDog := Animal(dog)
	animalCat, msgCat := Animal(cat)
	animalHam, msgHam := Animal(ham)
	animalTar, msgTar := Animal(tar)

	var amount float64
	if msgDog != "" {
		fmt.Println("Warning: ", msgDog)
	} else {
		amount = animalDog(5)
		fmt.Printf("La cantidad de alimento para %d perros es de: %.3f kg\n", 5, amount)
	}

	if msgCat != "" {
		fmt.Println("Warning: ", msgCat)
	} else {
		amount = animalCat(8)
		fmt.Printf("La cantidad de alimento para %d gatos es de: %.3f kg\n", 8, amount)
	}

	if msgHam != "" {
		fmt.Println("Warning: ", msgHam)
	} else {
		amount = animalHam(3)
		fmt.Printf("La cantidad de alimento para %d hamsters es de: %.3f kg\n", 3, amount)
	}

	if msgTar != "" {
		fmt.Println("Warning: ", msgTar)
	} else {
		amount = animalTar(6)
		fmt.Printf("La cantidad de alimento para %d tarantulas es de: %.3f kg\n", 6, amount)
	}
}
