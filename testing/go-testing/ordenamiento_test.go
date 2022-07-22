package calculadora

// Se importa el package testing y testify
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ejercicio 2 - Test Unitario Método Ordenar
// Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente,
// posteriormente diseñar un test unitario que valide el funcionamiento del mismo.
//  - Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.
//  - Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.
func TestOrdenamiento(t *testing.T) {
	// Datos de prueba para el testing
	num1, num2, num3, num4, num5, num6 := 4, 1, 2, 6, 3, 5

	// Resultado esperado en la prueba
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6}

	// Se testea la función
	resultado := Ordenamiento(num1, num2, num3, num4, num5, num6)

	// Se valida el resultado
	assert.Equal(t, resultadoEsperado, resultado, "Algunos elementos del slice no estan ordenados")
}
