package ordenamiento

import (
	"sort"
)

func OrdenamientoAsc(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	return numbers
}
