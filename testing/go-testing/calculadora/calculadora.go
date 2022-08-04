package calculadora

import (
	"errors"
	"sort"
)

func Sumar(n1, n2 int) int {
	return n1 + n2

}

func Restar(n1, n2 int) int {
	return n1 - n2
}

func Dividir(numerador, denominador int) (int, error) {
	if denominador == 0 {
		return 0, errors.New("El denominador no puede ser 0.")
	}

	return numerador / denominador, nil
}

func Sort(values ...int) []int {
	sort.Ints(values)
	return values
}
