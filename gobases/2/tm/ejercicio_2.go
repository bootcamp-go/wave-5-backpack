package main

import (
	"errors"
	"fmt"
)

func main() {
	c := []int{10, 8, 10, 7, 5, 8, 3, 4}
	p, err := promedio(c...)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Promedio: %.2f\n", p)
	}
}

func promedio(calificaciones ...int) (promedio float64, err error) {
	var total int
	for _, v := range calificaciones {
		if v < 0 {
			return 0, errors.New("No se admiten valores negativos")
		}

		total += v
	}

	return float64(total) / float64(len(calificaciones)), nil
}
