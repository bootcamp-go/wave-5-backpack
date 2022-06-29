package main

import "fmt"

func main() {
	fmt.Printf("El salario del empleado es de %f\n", getSalary(3000, "A"))
	fmt.Printf("El salario del empleado es de %f\n", getSalary(1500, "B"))
	fmt.Printf("El salario del empleado es de %f\n", getSalary(600, "C"))
}

func getSalary(minutes int, category string) float64 {
	var hours int = minutes / 60
	fmt.Println(hours)
	switch category {
	case "A":
		return (float64(hours) * 3000) * 1.5
	case "B":
		return (float64(hours) * 1500) * 1.2
	case "C":
		return float64(hours) * 1000
	}
	return 0
}
