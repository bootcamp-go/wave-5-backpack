package calculadora

import (
	"errors"
)

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("no se puede dividir por 0")
	}

	return num1 / num2, nil
}
