package testing

import "errors"

//Funcion que recibe dos enteros y retorna su suma
func Sum(num1, num2 int) int {
	return num1 + num2
}

//Funcion que recibe dos enteros y retorna su diferencia
func Res(num1, num2 int) int {
	return num1 - num2
}

//Funcion que recibe dos enteros y retorna su division
func Div(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("El denominador no puede ser 0")
	}

	return num1 / num2, nil
}
