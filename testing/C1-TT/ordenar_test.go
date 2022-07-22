package calculadora

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetodoInsercao(t *testing.T) {
	array := rand.Perm(10)
	result := MetodoInsercao(array)
	expectedResult := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}
