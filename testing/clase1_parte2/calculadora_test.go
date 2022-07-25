package calculadora

// Se importa el package testing
import "testing"

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8

	// Se ejecuta el test
	resultado := Sumar(num1, num2)

	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion suma() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}
}
