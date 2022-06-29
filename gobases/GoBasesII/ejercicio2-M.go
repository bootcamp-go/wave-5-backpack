package main

import (
	"errors"
	"fmt"
)

func main() {

	res, err := promedio(2, 3, 2, 2)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El promedio de notas es: ", res)
}

func promedio(notas ...float64) (float64, error) {

	var acu float64

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("El numero de esta posicion es negativo")
		}

		acu += value
	}

	return float64(acu) / float64(len(notas)), nil
}
