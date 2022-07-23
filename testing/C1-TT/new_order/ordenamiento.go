package neworder

import (
	"fmt"
	"sort"
)

func Ordenamiento(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	/*
		sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
		})
	*/

	for _, v := range s {
		fmt.Println(v)
	}
	return s
}
