package ordenamiento

func Ordenar(s []int) []int {
	ordenado := false
	o := []int{}

	for !ordenado {
		min := s[0]
		for _, v := range s {
			if min > v {
				min = v
			}
		}

		for i, v := range s {
			if min == v {
				if i != len(s)-1 {
					s = append(s[:i], s[i+1:]...)
				} else {
					s = s[:i]
				}
				break
			}
		}

		o = append(o, min)
		if len(s) == 0 {
			ordenado = true
		}
	}

	// metodo de libreria "sort"
	// sort.Slice(s, func(i, j int) bool {
	// 	return s[i] < s[j]
	// })

	return o
}
