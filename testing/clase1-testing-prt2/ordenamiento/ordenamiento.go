package calculadora

import "sort"

// Ordenamiento ...
func Ordenamiento(list_num []int) []int {
	sort.Slice(list_num, func(i, j int) bool {
		return list_num[j] < list_num[i]
	})
	return list_num
}
