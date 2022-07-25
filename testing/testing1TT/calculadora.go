package calculadora

import "fmt"

func Resta(n1, n2 int) int {
	res := n1 - n2
	return res
}

func Dividir(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("el denominador no puede ser 0")
	}
	res := n1 / n2
	return res, nil
}
