package ordenamiento

import "sort"

func Ordenar(list []int) []int {
	sort.Ints(list)
	return list
}
