package dividir

import "errors"

func Dividir(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New(ErrDivisionCero)
	}
	return a / b, nil
}
