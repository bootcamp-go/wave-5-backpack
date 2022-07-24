package calculadora

func ordenar(list []int) {
	for post := 1; post < len(list); post++ {
		key := list[post]
		j := post - 1
		//fmt.Println(j, key, list[j])
		for j >= 0 && key < list[j] {

			list[j+1] = list[j]
			//fmt.Println(list)
			j -= 1
		}
		list[j+1] = key
	}
}
