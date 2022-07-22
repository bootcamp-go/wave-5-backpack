package main

import "fmt"

const (
	minimun = "minimun"
	average = "average"
	maximun = "maximun"
)

func minValue(values ...int) float32 {
	min := 0
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return float32(min)
}

func maxValue(values ...int) float32 {
	max := 0
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return float32(max)
}

func avgValue(values ...int) float32 {
	count := 0
	for _, value := range values {
		count += value
	}
	return float32(count) / float32(len(values))
}

func operation(op string) func(values ...int) float32 {
	switch op {
	case "minimun":
		return minValue
	case "maximun":
		return maxValue
	case "average":
		return avgValue
	default:
		return nil
	}
}

func main() {
	minimun_op := operation(minimun)
	maximun_op := operation(maximun)
	average_op := operation(average)

	fmt.Printf("El minimo es %.2f\n", minimun_op(2, 3, 4, 6, -4))
	fmt.Printf("El maximo es %.2f\n", maximun_op(2, 3, 4, 6, -4))
	fmt.Printf("El promedio es %.2f\n", average_op(2, 3, 4, 6, -4))
}
