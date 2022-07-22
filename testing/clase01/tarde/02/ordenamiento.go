package ordenamiento

import (
	"sort"
)
func Ordenar(input []int) []int{ 
	sort.Ints(input)
	return input
}