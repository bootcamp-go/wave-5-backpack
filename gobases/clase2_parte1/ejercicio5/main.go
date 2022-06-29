package main

import (
	"errors"
	"fmt"
)

func main() {
	var kg float32 = 1000.00
	alimento, error := animal("tarantula")
	if error != nil {
		fmt.Println("Ocurrio un error, debido a que la mascota ingresada no existe")
	} else {
		cantComida := alimento(5)
		fmt.Println("La cantidad de comida que debe consumir el animal en gramos es: ", cantComida, "en KG es: ", cantComida/kg)
	}

}

const (
	tarantula = "tarantula"
	hamster   = "hamster"
	perro     = "perro"
	gato      = "gato"
)

func opTarantula(cantidad int) float32 {
	return float32(cantidad) * 150
}
func opPerro(cantidad int) float32 {
	return float32(cantidad) * 10000
}
func opGato(cantidad int) float32 {
	return float32(cantidad) * 50000
}
func opHamster(cantidad int) float32 {
	return float32(cantidad) * 250
}

func animal(mascota string) (func(cant int) float32, error) {
	switch mascota {
	case "tarantula":
		return opTarantula, nil
	case "hamster":
		return opHamster, nil
	case "perro":
		return opPerro, nil
	case "gato":
		return opGato, nil
	}
	return nil, errors.New("No existe el animal indicado")
}
