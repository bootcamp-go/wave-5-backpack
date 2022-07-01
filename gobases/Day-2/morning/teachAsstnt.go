package main

import "fmt"

const (
	minimum = "min"
	average = "avg"
	maximum = "max"
)

func getMin(notes ...float64) float64 {
	var min float64
	for i, note := range notes {
		if i == 0 {
			min = note
		}
		if note < min {
			min = note
		}
	}
	return min
}
func getMax(notes ...float64) float64 {
	var max float64
	for i, note := range notes {
		if i == 0 {
			max = note
		}
		if note > max {
			max = note
		}
	}
	return max
}
func getAvg(notes ...float64) float64 {
	var total float64
	for _, note := range notes {
		total += note
	}
	return total / float64(len(notes))
}

func operation(noteType string) func(notes ...float64) float64 {
	switch noteType {
	case minimum:
		return getMin
	case average:
		return getAvg
	case maximum:
		return getMax
	}
	return nil
}

func main() {
	op := operation(maximum)
	note := op(5, 4, 1, 3, 8, 4, 5)
	fmt.Println(note)
}
