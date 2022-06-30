package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float64) (float64, error) {
	var sumaNotas float64
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Una nota es menor a cero")
		}
		sumaNotas += nota
	}

	return (sumaNotas / float64(len(notas))), nil
}

func main() {
	promedio, err := calcularPromedio(-1, 3, 5, 5)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("El promedio es :", promedio)
	}
}
