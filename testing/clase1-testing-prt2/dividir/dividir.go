package dividir

import "errors"

// Dividir ... Función que recibe dos enteros (numerador y denominador) y retorna la división resultante
func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("El denominador debe ser diferente de 0 *cero*")
	}
	return num / den, nil
}
