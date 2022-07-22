package ordenar

import (
	"sort"
)

func OrdenasAsc(sliceNum []int) []int {
	sort.Slice(sliceNum, func(i, j int) bool {
		return sliceNum[i] < sliceNum[j]
	})
	return sliceNum
}
