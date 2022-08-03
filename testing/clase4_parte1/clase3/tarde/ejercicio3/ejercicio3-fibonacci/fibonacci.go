package fibonacci

import (
	"fmt"
	"strings"
)

func Fibonnaci(num int) (string, int) {
	var result []int
	var total int

	for i := 0; i < num; i++ {
		if i == 0 {
			result = append(result, i)
			total += i
		}

		if i == 1 {
			result = append(result, i)
			total += i
		}

		if i >= 2 {
			result = append(result, (result[i-1] + result[i-2]))
			total += result[i-1] + result[i-2]
		}
	}

	sucession := strings.Fields(fmt.Sprint(result))
	return strings.Trim(strings.Join(sucession, " "), "[]"), total
}
