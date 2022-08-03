package calculadora

import "testing"

func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 3
	resultadoEsperado := 2

	resultado := Restar(num1, num2)

	if resultadoEsperado != resultado {
		t.Errorf("Function restar() arrojo el resultado = %v, pero el esperado es = %v", resultado, resultadoEsperado)
	}
}
