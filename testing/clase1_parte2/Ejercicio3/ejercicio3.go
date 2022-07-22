package ejercicio3

import "errors"

func dividir(num, den int) (int, error) {
	if den <= 0 {
		return 0, errors.New("error: el denominador no puede ser 0 o negativo")
	}
	return num / den, nil
}
