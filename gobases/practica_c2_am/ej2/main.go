package main

import (
	"errors"
	"fmt"
)

func main() {
	prom, err := promedio(33, 45, 60, 70, 55)
	fmt.Println("Calificaciones: ", 33, 45, 60, 70, 55)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio del curso es: ", prom)
	}

	prom, err = promedio(33, -45, 60, 70, 55)
	fmt.Println("Calificaciones: ", 33, -45, 60, 70, 55)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio del curso es: ", prom)
	}
}

func promedio(nEnteros ...int) (int, error) {
	suma := 0
	for _, n := range nEnteros {
		if n < 0 {
			return 0, errors.New("No puede ingresar notas negativas!")
		}
		suma += n
	}
	return suma / len(nEnteros), nil
}
