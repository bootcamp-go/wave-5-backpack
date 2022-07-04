package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func aDog(value float64) float64 {
	amountD := value * 10
	return amountD

}

func aCat(value float64) float64 {
	amountC := value * 5
	return amountC

}

func aHamster(value float64) float64 {
	amountH := value * 0.25
	return amountH

}

func aTarantula(value float64) float64 {
	amountT := value * 0.15
	return amountT

}

func animal(especieAnimal string) (func(value float64) float64, error) {

	switch especieAnimal {
	case "dog":
		return aDog, nil
	case "cat":
		return aCat, nil
	case "hamster":
		return aHamster, nil
	case "tarantula":
		return aTarantula, nil
	}

	return nil, errors.New("El animal no existe")
}

func main() {
	opera, msg := animal("cat")
	if msg != nil {
		fmt.Println(msg)
	} else {
		r := opera(2)
		fmt.Printf("La cantidad de comida es: %v kg\n", r)
	}

}
