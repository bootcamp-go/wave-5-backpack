package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("el divisor no puede ser cero")
	}

	return dividendo / divisor, nil
}

func main() {
	res, err := division(2, 0)

	if err != nil {
		fmt.Printf("Ocurrio un error: %v \n", err)
	}
	fmt.Printf("Resultado %.2f \n", res)
}
