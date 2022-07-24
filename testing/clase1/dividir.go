package calculadora

import "fmt"

func dividir(a, b int) (int, error) {
	if b == 0 {
		return -1, fmt.Errorf("el demoninador no puede ser 0")
	}
	return int(a / b), nil
}
