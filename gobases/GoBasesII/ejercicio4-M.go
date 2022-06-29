package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {

	minFunc := operation(minimum)
	averageFunc := operation(average)
	maxFunc := operation(maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue)
	fmt.Println(averageValue)
	fmt.Println(maxValue)

}

func minFunc(n ...float64) float64 {

	min := n[0]

	for i := 0; i < len(n); i++ {
		if n[i] < min {
			min = n[i]
		}
	}

	return min
}

func maxFunc(n ...float64) float64 {

	max := n[0]

	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}

	return max
}

func averageFunc(n ...float64) float64 {

	var cont float64

	for _, value := range n {
		cont += value
	}

	return float64(cont) / float64(len(n))
}

func operation(operacion string) func(n ...float64) float64 {
	switch operacion {
	case minimum:
		return minFunc
	case maximum:
		return maxFunc
	case average:
		return averageFunc
	default:
		fmt.Println("Calculo no definido")
	}
	return nil
}
