package ordenamiento

import "sort"

func Ordenar(s []int) []int {
	sort.Sort(sort.IntSlice(s))
	return s
}
