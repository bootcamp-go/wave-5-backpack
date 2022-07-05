package main

import (
	"errors"
	"fmt"
)

func sumaNotas(params ...float64) (float64, error) {
	var sum float64
	for _, num := range params {
		if num < 0 {
			return 0, errors.New("no pueden haber numeros negativos")
		}
		sum += num
	}
	return sum, nil
}

func promediarNotas(notas ...float64) (float64, error) {
	suma, errSuma := sumaNotas(notas...)
	return suma / float64(len(notas)), errSuma
}

func main() {
	promedio, err := promediarNotas(5, 0, 5)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(promedio)
}
