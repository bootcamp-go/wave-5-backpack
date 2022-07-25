package C1TT

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

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 5
	num2 := 3
	resultadoEsperado := 2

	// Se ejecuta el test
	resultado := Restar(num1, num2)

	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
	}
}
