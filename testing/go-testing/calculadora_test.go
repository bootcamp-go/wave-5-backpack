package calculadora

// Se importa el package testing
import (
	"testing"
)

// Ejercicio 1 - Test Unitario Restar
// Para el método Restar() visto en la clase, realizar el test unitario correspondiente. Para eso:
//  - Dentro de la carpeta go-testing crear un archivo calculadora.go con la función a probar.
//  - Dentro de la carpeta go-testing crear un archivo calculadora_test.go con el test diseñado.
func TestRestar(t *testing.T) {
	// Datos de prueba para el testing
	num1 := 8
	num2 := 5
	// Resultado esperado en la prueba
	resultadoEsperado := 3

	// Se testea la función
	resultado := Restar(num1, num2)

	// Se valida el resultado
	if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() devolvio = %v, valor esperado = %v\n", resultado, resultadoEsperado)
	}
}
