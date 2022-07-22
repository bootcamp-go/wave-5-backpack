package main

import "fmt"

func main() {
	var salary int = 75000
	fmt.Println(taxes(salary))
}

func taxes(salary int) float64 {
	switch {
	case salary > 50000:
		return float64(salary) * 0.17
	case salary > 150000:
		return float64(salary) * 0.1
	default:
		return float64(salary)
	}
}
