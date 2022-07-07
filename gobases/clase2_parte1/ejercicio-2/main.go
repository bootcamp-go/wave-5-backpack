package main

import (
	"errors"
	"fmt"
)

func main() {
	prom, err := promedio(1, 2, 3, 4, 5)
	fmt.Println("Números: ", 1, 2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", prom)
	}

	prom, err = promedio(1, 2, -3, 4, 5)
	fmt.Println("Números: ", 1, 2, -3, 4, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", prom)
	}
}

func promedio(nEnteros ...int) (int, error) {
	suma := 0
	for _, n := range nEnteros {
		if n < 0 {
			return 0, errors.New("Ingrese solo enteros positivos")
		}
		suma += n
	}
	return suma / len(nEnteros), nil
}
