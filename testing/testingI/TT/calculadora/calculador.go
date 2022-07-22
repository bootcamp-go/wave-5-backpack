package calculadora

import "errors"

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("el divisor no puede ser 0")
	}
	return num1 / num2, nil
}
