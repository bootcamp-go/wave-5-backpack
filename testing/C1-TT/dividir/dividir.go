package dividir

import (
	"fmt"
)

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return den, fmt.Errorf("El denominador no puede ser cero")
	}

	return num / den, nil
}
