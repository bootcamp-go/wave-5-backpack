package main

import "fmt"

func giveSalaryByType(salaryType string, minutesWorked int) float64 {
	// Map statement
	var jobType = make(map[string][]float64)
	jobType["A"] = []float64{1000.00}
	jobType["B"] = []float64{1500.00, .20}
	jobType["C"] = []float64{3000.00, .50}
	fmt.Println(jobType)
	salary := jobType[salaryType][0] * (float64(minutesWorked) / 60)
	if salaryType == "B" || salaryType == "C" {
		bonus := salary * jobType[salaryType][1]
		return bonus + salary
	}
	return salary
}

func main() {
	fmt.Println(giveSalaryByType("B", 60))
}
