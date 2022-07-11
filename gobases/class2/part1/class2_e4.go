package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(values ...int) int {
	result := 0
	for _, value := range values {
		if result == 0 {
			result = value
		} else if result > value {
			result = value
		}
	}
	if result == 0 {
		return result
	} else {
		return result
	}

}

func averageFunc(values ...int) int {
	var result = 0
	ran := 0
	for _, value := range values {
		ran++
		if value < 0 {
			return 0
		}
		result += value
	}

	if ran == 0 {
		return result
	} else {
		return result / ran
	}
}

func maxFunc(values ...int) int {
	result := 0
	for _, value := range values {
		if result == 0 {
			result = value
		} else if result < value {
			result = value
		}
	}
	if result == 0 {
		return result
	} else {
		return result
	}

}

func operations(values []int, operation func(value ...int) int) int {
	var result int
	for i, value := range values {
		if i == 0 {
			result = value
		} else {
			result = operation(values...)
		}
	}

	return result
}

func operationRequest(operator string, values ...int) int {
	switch operator {
	case minimum:
		return operations(values, minFunc)
	case average:
		return operations(values, averageFunc)
	case maximum:
		return operations(values, maxFunc)
	}
	return 0
}

//  minFunc, err := operation(minimum)
//  averageFunc, err := operation(average)
//  maxFunc, err := operation(maximum)

//  minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
//  averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
//  maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

func main() {
	fmt.Println(operationRequest(minimum, 4, 6, 8, 2, 4))
	fmt.Println(operationRequest(average, 4, 6, 8, 2, 4))
	fmt.Println(operationRequest(maximum, 4, 6, 8, 2, 4))
}
