package calculadora

import "fmt"

// Restar

func Restar(num1, num2 int) int {
	return num1 - num2
}

//Dividir

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("el denominador no puede ser 0")
	}
	return num / den, nil
}
