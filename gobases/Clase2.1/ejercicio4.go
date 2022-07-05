package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func calcMinimum(grades []float64) float64 {
	min := -1.0
	for _, grade := range grades {
		if min == -1 {
			min = grade
		}
		if grade < min {
			min = grade
		}
	}
	return min
}

func calcAverage(grades []float64) float64 {
	avg := 0.0
	for _, grade := range grades {
		avg = avg + grade
	}
	return avg / float64(len(grades))
}

func calcMax(grades []float64) float64 {
	max := 0.0
	for _, grade := range grades {
		if grade > max {
			max = grade
		}
	}
	return max
}

func operationsSelector(operation string) (func(grades []float64) float64, error) {
	switch operation {
	case minimum:
		return calcMinimum, nil
	case maximum:
		return calcMax, nil
	case average:
		return calcAverage, nil
	}
	return nil, errors.New("Invalid operation")
}

func calculate(operation string, grades ...float64) float64 {
	oper, _ := operationsSelector(operation)
	return oper(grades)
}

func main() {
	fmt.Println(calculate(average, 1.0, 2.0, 3.0))
}
