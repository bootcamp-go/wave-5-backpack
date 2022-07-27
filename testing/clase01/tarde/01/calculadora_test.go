package calculadora

import (
    "testing"
    "github.com/stretchr/testify/assert" // Se importa testify
)
func TestRestar(t *testing.T) {
    // Se inicializan los datos a usar en el test (input/output)
    num1 := 5
    num2 := 3
    resultadoEsperado := 2
    // Se ejecuta el test
    resultado := Restar(num1, num2)
    // Se validan los resultados
    assert.NotNil(t, resultado) // se valida que no sea nulo
    assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}