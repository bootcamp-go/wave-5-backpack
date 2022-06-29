package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(calcularPromedio(3, 6, 9))
	fmt.Println(calcularPromedio(3, -1, 9))
}

func calcularPromedio(notas ...int) (float64, error) {
	var resultado float64
	var i float64
	for _, nota := range notas {
		if nota < 0 {
			return -1, errors.New("existe una nota negativa, devolviendo -1 como resultado")
		} else {
			resultado += float64(nota)
			i++
		}
	}
	return resultado / i, nil
}
