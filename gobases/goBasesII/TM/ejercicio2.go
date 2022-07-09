package main

import (
	"errors"
	"fmt"
)

func calPromedio(values ...int) (float64, error) {
	var resultado float64
	for _, value := range values {
		if value < 0 {
			return 0, errors.New("Hay calificaciones con valores negativos, no es posible retornar un resultado")
		}
		resultado += float64(value)
	}
	resultado = resultado / float64(len(values))
	return resultado, nil
}

func main() {
	prom, err := calPromedio(1, -2, 3, 4, 5, 6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("resultado: %.2f\n", prom)
	}

}
