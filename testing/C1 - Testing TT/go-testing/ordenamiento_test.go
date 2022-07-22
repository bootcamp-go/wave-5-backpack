package ejercicio

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	slice := []int{5, 3, 4, 7, 8, 9}
	resultadoEsperado := []int{3, 4, 5, 7, 8, 9}

	// Se ejecuta el test
	resultadoOrdenar := Ordenar(slice)

	// Se validan los resultados del ordenamiento
	assert.Equal(t, resultadoEsperado, resultadoOrdenar, "El slice esperado en el ordenamiento no coincide con el slice obtenido en la función Ordenar()")

}
