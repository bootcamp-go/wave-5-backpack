package main

import (
	"errors"
	"fmt"
)

func promedio(valores ...float64) (float64, error) {
	a, b, c := 0, 0, len(valores)

	for i, valor := range valores {

		if valor < 0 {
			a++
			i++
		} else {
			b += int(valor)
		}

	}

	if a != 0 {
		return 0, errors.New("No pueda haber número negativo")
	} else {
		d := float64(b)
		e := float64(c)
		return d / e, nil
	}

}
func main() {

	res, err := promedio(5, 3, -4)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Promedio:", res)
	}

}
