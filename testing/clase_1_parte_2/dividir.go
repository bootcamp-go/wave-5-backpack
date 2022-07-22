package clase_1_parte_2

import "errors"

var (
	ErrDominadorCero = errors.New("el denominador no puede ser 0")
)

func Dividir(num, dom int) (int, error) {
	if dom == 0 {
		return 0, ErrDominadorCero
	}
	return num / dom, nil
}
