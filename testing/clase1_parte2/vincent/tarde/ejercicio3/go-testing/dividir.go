package dividir

import "fmt"

func Dividir(num, dem int) (int, error) {
	if dem <= 0 {
		return 0, fmt.Errorf("el denominador no puede ser 0")
	}

	return num / dem, nil
}
