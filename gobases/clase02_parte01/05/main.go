/*
Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/

package main
import (
	"fmt"
	"errors"
)

const (
	perro = "dog"
	gato = "cat"
	tarantula = "tarantula"
	hamster = "hamster"
)

func alimPerro(cant float64) float64 {
	return 10 * cant
}

func alimGato(cant float64) float64 {
	return 5 * cant
}

func alimTarantula(cant float64) float64 {
	return 0.150 * cant
}

func alimHamster(cant float64) float64 {
	return 0.250 * cant
}


func Animal (animal string) (func(cant float64) float64, error){
	switch animal {
	case perro:
		return alimPerro , nil
	case gato:
		return alimGato , nil
	case tarantula:
		return alimTarantula , nil
	case hamster:
		return alimHamster , nil
	}
	return nil, errors.New("No existe ese animal")
}

func main (){
	animalDog, _ := Animal(perro)
	animalCat, _ := Animal(gato)

	var amount float64
	amount+= animalDog(2)
	amount+= animalCat(1)

	fmt.Println(amount)
}

