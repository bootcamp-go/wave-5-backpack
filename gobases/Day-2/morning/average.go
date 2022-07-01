package main

import "fmt"

func averageCompute(notes ...float64) float64 {
	var sum float64
	for _, note := range notes {
		sum += note
	}
	return sum / float64(len(notes))
}

func main() {
	fmt.Println(averageCompute(10, 5, 3))
}
