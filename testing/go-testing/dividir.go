package main

import "errors"

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("El denominador no puede ser 0")
	}
	return num / den, nil
}
