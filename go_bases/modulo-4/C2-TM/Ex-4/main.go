package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

}

func multiFunction(op string) (func(i ...int) float64, error) {
	switch op {
	case "minimum":
		return Opmin, nil
	case "average":
		return Opaverage, nil
	case "maximum":
		return Opmax, nil
	default:
		return AnonError, errors.New("Esa operacion no existe")
	}
}

func Opmin(i ...int) float64 {
	var result float64

	sort.Ints(i)
	result = float64(i[0])
	fmt.Println(result)
	return result
}
func Opaverage(i ...int) float64 {
	var result float64

	sort.Ints(i)
	result = float64(i[0])
	fmt.Println(result)
	return result
}
func Opmax(i ...int) float64 {
	var result float64

	sort.Ints(i)
	result = float64(i[len(i)])
	fmt.Println(result)
	return result
}
func AnonError(i ...int) float64 {
	var result float64
	result = float64(0)
	return result
}
