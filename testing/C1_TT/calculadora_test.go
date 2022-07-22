package C1_TT

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestar(t *testing.T) {
	resultadoEsperado := 10
	resultado := Restar(30, 20)

	//usando testing
	if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() arrojo el resultado: %v, pero el resultado esperado es: %v", resultado, resultadoEsperado)
	}

	//usando testify
	assert.Equal(t, resultadoEsperado, resultado, "El resultado esperado es %v", resultadoEsperado)
}
