package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.
//  - Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
//  - Dentro de la carpeta go-testing crear un archivo dividir_test.go con el test diseñado.
func TestDividir(t *testing.T) {
	// Datos de prueba de éxito
	num1 := 15
	num2 := 3
	resultadoEsperado := 5

	// Se testea la función
	resultado1, err := Dividir(num1, num2)

	// Se valida el resultado
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado1, "La division de %v / %v no es %v", num1, num2, resultado1)

	// Datos de prueba de error
	num1 = 18
	num2 = 0
	resultadoEsperado = 0

	// Se testea la función
	resultado2, err2 := Dividir(num1, num2)

	if err2 != nil {
		assert.Nil(t, err2.Error())
	}
	// Se valida el resultado
	assert.Equal(t, resultadoEsperado, resultado2, "Deben ser iguales")
}
