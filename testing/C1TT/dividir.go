package C1TT

import "fmt"

func NewDividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return num1 / num2, nil
}
