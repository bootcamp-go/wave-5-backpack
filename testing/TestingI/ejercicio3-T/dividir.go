package dividir

import "errors"

func Dividir(num1, num2 int) (int, error) {

	if num2 == 0 {
		return 0, errors.New("el denominador no puede ser cero")
	}

	return num1 / num2, nil
}
