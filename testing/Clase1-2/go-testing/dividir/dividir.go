package dividir

import (
	"errors"
)

func Dividir(num, den int) (int, error) {
	if den == 0 {
		err := errors.New("El denominador no puede ser 0")
		return 0, err
	}
	return (num / den), nil
}
