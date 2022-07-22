package calculadora

import "errors"

func Dividir(num1, num2 int) (int, error) {

	if num2 == 0 {
		return 0, errors.New("No se puede dividir por 0")
	}
	return num1 / num2, nil
}
