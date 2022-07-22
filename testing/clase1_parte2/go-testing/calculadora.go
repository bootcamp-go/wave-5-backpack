package claseTesting

import "errors"

func Restar(a, b int) int {
	return a - b
}

func Dividir(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("El denominador no puede ser 0")
	}
	return float64(a) / float64(b), nil
}
