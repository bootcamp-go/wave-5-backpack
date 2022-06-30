package main

import (
	"errors"
	"fmt"
)

func calcularPromedios(notas ...int) (float64, error) {
	var resultado float64

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("error")
		}
		resultado += float64(nota)
	}

	resultado = resultado / float64(len(notas))
	return resultado, nil
}

func main() {
	res, err := calcularPromedios(-1, 72, 5, 13)
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Printf("%.2f \n", res)
	}
}
