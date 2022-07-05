/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #5:  Calcular cantidad de alimento
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		An animal shelter needs to figure out how much food to buy for pets.
 		At the moment they only have tarantulas, hamsters, dogs, and cats, but
		it is expected that there could be many more animals to shelter.
			1.	dog need 10 kg of food
			2.	cat 5 kg
			3.	Hamster 250 grams.
			4.Tarantula 150 grams

		You are requested to:
 			1.	implement a function Animal that receives as parameter a text
				type value with the specified animal and that returns a function
				and a message (in case the animal does not exist).
 			2.	A function for each animal that calculates the amount of food
				based on the amount of the specified animal type.

		Example:
		 const  (
			dog  =  "dog"
			cat  =  "cat"
		)
		...
		animalDog  ,  msg  :=  Animal  (dog)
		animalCat  ,  msg  :=  Animal  (cat)
		...
		var  amount  float64
		amount+=  animaldog  (  5  )
		amount+=  animalCat  (  8  )

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARIES
package main

import (
	"errors"
	"fmt"
)

// CONSTANTS
const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

// FUNCTIONS
func feedDog(alimento float64) float64 {
	return alimento * 10 // Son 10 kg por cada perro
}

func feedCat(alimento float64) float64 {
	return alimento * 5 // Son 5 kg por cada gato
}

func feedHamster(alimento float64) float64 {
	return alimento * 250 // Son 250 gr por cada hamster
}

func feedTarantula(alimento float64) float64 {
	return alimento * 150 // Son 150 gr por cada tarantula
}

func Animal(elecAnimal string) (func(alimento float64) float64, error) {
	switch elecAnimal {
	case dog:
		return feedDog, nil
	case cat:
		return feedCat, nil
	case hamster:
		return feedHamster, nil
	case tarantula:
		return feedTarantula, nil
	}
	return nil, errors.New("** Esta operacion no se encuentra en la lista **")
}

//	MAIN PROGRAM
func main() {
	fmt.Println("\n\t|| Calcular cantidad de Alimento ||")

	var amount float64

	// Dog
	animalDog, msg := Animal(dog)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalDog(5)
		fmt.Printf(">  Cantidad a obtener para perros: %.2f[kg]\n", amount)
	}

	// Cat
	animalCat, msg := Animal(cat)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalCat(8)
		fmt.Printf(">  Cantidad a obtener para gatos: %.2f[kg]\n", amount)
	}

	// Hamster
	animalHamster, msg := Animal(hamster)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalHamster(3)
		fmt.Printf(">  Cantidad a obtener para hamsters: %.2f [gr]\n", amount)
	}

	// Tarantula
	animalTarantula, msg := Animal(tarantula)
	if msg != nil {
		fmt.Println(msg)
	} else {
		amount += animalTarantula(4)
		fmt.Printf(">  Cantidad a obtener para tarantulas: %.2f [gr]\n", amount)
	}
}
