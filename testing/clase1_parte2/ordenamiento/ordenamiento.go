package ordenamiento

import "sort"

func Ordenar(values []int) []int {
	sortedSlice := values
	sort.Ints(sortedSlice)
	return sortedSlice
}