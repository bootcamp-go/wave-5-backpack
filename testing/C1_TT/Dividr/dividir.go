package Dividr

import "errors"

func Dividir(num1, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("el denominador no puede ser 0")
	}
	return num1 / den, nil
}
