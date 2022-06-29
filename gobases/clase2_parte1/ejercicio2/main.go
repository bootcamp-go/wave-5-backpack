package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := promedio(9, 2, 7.5, 10)
	if err == nil {
		fmt.Println("El promedio es", res)
	} else {
		fmt.Println("Ocurrio un error:", err)
	}

}

func promedio(val ...float32) (float32, error) {
	var suma float32
	var cant int
	for _, valor := range val {
		if valor < 0 {
			return 0, errors.New("Algunas de las notas es un numero negativo")
		} else {
			suma += valor
			cant += 1
		}
	}
	return suma / float32(cant), nil
}
