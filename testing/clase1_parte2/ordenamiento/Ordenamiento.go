package ordenamiento

import (
	"sort"
)

func Ordenar(nums []int) []int {
	sort.Ints(nums)
	return nums
}
