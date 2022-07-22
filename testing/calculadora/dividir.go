package calculadora

import "fmt"

func Dividir(num, den int) (int, error) {
	if den == 0 {
		err := fmt.Errorf("el denominador no puede ser cero")
		return 0, err
	}
	return num / den, nil
}
