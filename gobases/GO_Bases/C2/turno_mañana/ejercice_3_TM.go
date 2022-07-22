package main

import "fmt"

func minToHours(min int) float32 {
	return float32(min) / 60
}

func calculated_salary(category string, minutes int) float32 {
	switch category {
	case "A":
		return 3000 * minToHours(minutes) * 1.5
	case "B":
		return 1500 * minToHours(minutes) * 1.2
	case "C":
		return 3000 * minToHours(minutes)
	default:
		return 0
	}
}

func main() {
	min, category := 15000, "B"
	fmt.Printf("Salario: %.2f\n", calculated_salary(category, min))
}
