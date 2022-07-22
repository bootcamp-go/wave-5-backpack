package calculadora

import (
	"errors"
)

func Sumar(num1, num2 int) int {
	return num1 + num2
}

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num1, num2 int) (int, error) {

	if num2 == 0 {
		return 0, errors.New("no se puede dividir por cero")
	}

	return num1 / num2, nil
}

func Multiplicar(num1, num2 int) int {
	return num1 * num2
}
