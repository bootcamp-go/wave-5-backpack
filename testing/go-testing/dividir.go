package calculadora

import "errors"

// Función para dividir 2 enteros
// Cambiar el método para que no sólo retorne un entero sino también un error.
// Incorporar una validación en la que si el denominador es igual a 0,
// retorna un error cuyo mensaje sea "El denominador no puede ser 0"
func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("el denominador no puede ser 0")
	}
	return num / den, nil
}
