package testing

import (
	"errors"
	"sort"
)

func Restar(n1, n2 int) int {
	return n1 - n2
}

func Ordenar(m []int) []int {

	sort.Ints(m)

	return (m)
}

func Dividir(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("El denominador no puede ser 0")
	}
	return (n1 / n2), nil

}
