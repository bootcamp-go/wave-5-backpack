package ejercicio

import (
	"sort"
)

//Tomamos el slice y hacemos el sort
func Ordenar(slice []int) []int {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}
