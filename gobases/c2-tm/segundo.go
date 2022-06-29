package main

import (
	"errors"
	"fmt"
)

func promedio(values ...float64) (float64, error) {
	var resultado float64
	for _, value := range values {
		if value < 0 {
			return 0, errors.New("La calificacion no puede ser negativa")
		}
		resultado += value
	}

	return (resultado / float64(len(values))), nil
}

func main() {
	res, err := promedio(5, 4, 3, 2.8, 4, 5, 3, 2, 4)

	if err != nil {
		fmt.Println("La calificacion no puede ser negativa")
	} else {
		fmt.Printf("%.2f\n", res)
	}
}
