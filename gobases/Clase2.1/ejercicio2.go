package main

import (
	"errors"
	"fmt"
)

func calcPromedio(notas ...int) (int, error) {
	var promedio int = 0
	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("Nota negativa")
		}
		promedio = promedio + nota
	}
	return promedio / len(notas), nil
}

func main() {
	prom, err := calcPromedio(3, 3, 3, 3, -1)
	fmt.Println(prom)
	fmt.Println(err)
}
