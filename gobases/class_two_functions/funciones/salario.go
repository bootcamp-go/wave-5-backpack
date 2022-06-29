package main

import (
	"fmt"
	"strings"
)

const (
	totalOfHoursAtMonth int = 240
	salaryC             int = 1000
	salaryB             int = 1500
	salaryA             int = 3000
	percentMonthB       int = 20
	percentMonthA       int = 50
	minutesPerHour      int = 60
)

func main() {

	totalOfMinutesWorking, employeeCategory := getData()
	salary := calculateSalary(totalOfMinutesWorking, employeeCategory)

	fmt.Println("El salario a pagarle al empleado es:", salary)
}

func getData() (int, string) {
	var totalOfMinutesWorking int = 0
	var employeeCategory string
	var valid bool = false
	for !valid {
		fmt.Println("------Ingrese el total de minutos trabajados por el empleado")
		fmt.Scanf("%v", &totalOfMinutesWorking)
		fmt.Println("------Ingrese la categoría del empleado")
		fmt.Scanf("%s", &employeeCategory)
		valid = validateData(totalOfMinutesWorking, employeeCategory)
		if !valid {
			fmt.Println("¡ups! algun dato es incorrecto, intente de nuevo.")
		}
	}

	return totalOfMinutesWorking, employeeCategory
}

func validateData(minutes int, employeeCategory string) bool {

	if minutes > 0 && ((strings.Compare(employeeCategory, "A") == 0) ||
		(strings.Compare(employeeCategory, "B") == 0) ||
		(strings.Compare(employeeCategory, "C") == 0)) {
		return true
	}
	return false
}
func calculateSalary(minutes int, category string) float32 {
	var totalOfHours float32 = float32(minutes) / float32(minutesPerHour)
	var amountOfSalaryPerMonth float32 = 0
	var amountOfSalaryPerHours float32 = 0

	switch category {
	case string(rune('A')):
		amountOfSalaryPerHours = float32(salaryA) * (totalOfHours)                     // 18 000
		amountOfSalaryPerMonth = (float32(salaryA) * float32(totalOfHoursAtMonth)) / 2 //720 000
	case string(rune('B')):
		amountOfSalaryPerHours = float32(salaryB) * (totalOfHours)                     // 9 000
		amountOfSalaryPerMonth = (float32(salaryB) * float32(totalOfHoursAtMonth)) / 5 // 72000
	case string(rune('C')):
		amountOfSalaryPerHours = float32(salaryC) * (totalOfHours) // 9 000
		amountOfSalaryPerMonth = 0
	}

	return amountOfSalaryPerMonth + amountOfSalaryPerHours
}
