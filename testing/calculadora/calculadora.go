package calculadora

import "errors"

//Funcion que recibe dos enteros y retorna la suma resultante

func Sumar(num1, num2 int) int {
	return num1 + num2
}

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("El denominador no puede ser 0")
	}
	return num / den, nil
}
