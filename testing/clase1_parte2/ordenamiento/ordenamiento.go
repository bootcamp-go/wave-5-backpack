package ordenamiento

import "sort"

func Ordenar(numSlice []int) []int {
	sort.Slice(numSlice, func(i, j int) bool {
		return numSlice[i] < numSlice[j]
	})
	return numSlice
}
