package calculadora

import "fmt"

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}

	if den < 0 {
		return 0, fmt.Errorf("El denominador no puede ser negativo")
	}
    return num / den, nil
}
