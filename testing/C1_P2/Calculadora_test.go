package C1_P2

import "testing"

func TestRestar(t *testing.T) {
	num1 := 30
	num2 := 15
	expected := 15
	result := Restar(num1, num2)

	if result != expected {
		t.Errorf("El resultado %v no es el esperado, debería ser %v", result, expected)
	}

}
