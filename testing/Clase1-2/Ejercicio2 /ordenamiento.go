package ordenamiento

import (
	"fmt"
	"sort"
)

func Ordenar(desorden []int) []int {
	fmt.Println(desorden)
	sort.Ints(desorden)
	fmt.Println(desorden)
	return desorden
}
