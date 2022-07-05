package main

import "fmt"

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(animal string) func(cantidad int) float64 {
	switch animal {
	case perro:
		return fperro
	case gato:
		return fgato
	case hamster:
		return fhamster
	case tarantula:
		return ftarantula
	}
	return nil
}

//funciones animales
func fperro(cantidad int) float64 {
	return float64(cantidad * 10)
}

func fgato(cantidad int) float64 {
	return float64(cantidad * 5)
}

func fhamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}

func ftarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}

func main() {
	oper := Animal("tarantula")
	if oper != nil {
		r := oper(5)
		fmt.Printf("La comida del animal es: %.1f kg\n", r)
	} else {
		fmt.Println("El animal no existe")
	}
}
