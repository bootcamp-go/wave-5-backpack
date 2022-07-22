package main

func AscendingSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		var minIdx = i
		for j := i; j < len(list); j++ {
			if list[j] < list[minIdx] {
				minIdx = j
			}
		}
		list[i], list[minIdx] = list[minIdx], list[i]
	}
	return list
}
