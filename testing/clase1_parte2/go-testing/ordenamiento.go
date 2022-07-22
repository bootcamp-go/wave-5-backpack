package calculadora
import (
	"sort"
)

func Ordenamiento(l []int) []int {
	sort.Slice(l, func(i, j int) bool {
		return l[j] > l[i]
	 })
	return l
}
