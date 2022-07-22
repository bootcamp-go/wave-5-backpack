package testing_go

import "errors"

func Dividir(num1, den int) (int, error) {

	if den == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}
	return num1 / den, nil
}
