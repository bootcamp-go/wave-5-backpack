package main

import (
	"errors"
	"fmt"
)

var promedio int

func calcularPromedio(notas ...int) (int, error) {
	var sumaNotas int
	numNotas := len(notas)
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("No se admiten valores negativos")
		}
		sumaNotas += nota
	}
	if numNotas == 0 {
		return 0, errors.New("No hay notas")
	} else {
		return sumaNotas / numNotas, nil
	}

}

func main() {
	fmt.Println(calcularPromedio(5, 4, 3, 4, 5, 6))
}
