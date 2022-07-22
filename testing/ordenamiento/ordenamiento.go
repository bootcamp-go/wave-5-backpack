package ordenamiento

import "sort"

func OrderSlice(dataInt []int) []int {
	sort.Ints(dataInt)
	return dataInt
}
