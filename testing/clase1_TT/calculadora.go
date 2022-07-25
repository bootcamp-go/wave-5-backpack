package calculadora

import "fmt"

func Restar(a, b int) int {
	return a - b
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("the denominator should be different of 0")
	}
	return num / den, nil
}

func OrderAsc(slice []int) []int {
	var indexI, indexJ, min, val int
	var order bool
	for i, v := range slice {
		min, val = v, slice[i]
		indexI = i
		for j := i; j < len(slice); j++ { // runing from the last position and the last min found
			if min > slice[j] {
				min = slice[j] // minimum value , at i position of slice a
				indexJ = j
				order = true
			}
		}
		if order {
			slice[indexI] = min
			slice[indexJ] = val
			order = false
		}
	}
	return slice
}
