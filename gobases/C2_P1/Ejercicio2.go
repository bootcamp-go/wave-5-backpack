package main

import (
	"errors"
	"fmt"
)

func promedio(valores ...int) (int, error) {
	var result int
	for _, suma := range valores {
		if suma < 0 {
			return 0, errors.New("Negative values not acepted")
		}
		result += suma
	}

	return result / len(valores), nil
}

func main() {
	oper, err := promedio(8, 9, 10, 4, 6, 7, 8, 8, 10)
	if err != nil {
		fmt.Println("Error presentado:", err)
	} else {
		fmt.Println("El promedio de notas es", oper)
	}
}
