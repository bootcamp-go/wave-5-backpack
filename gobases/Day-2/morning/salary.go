package main

import "fmt"

func taxesSalary(income float64) {
	var taxe float64
	if income >= 150000.00 {
		taxe = income * 0.27
		fmt.Println("Taxe is 27%")
		fmt.Println(income)
		fmt.Println(taxe)
	} else if income >= 50000.00 {
		taxe = income * .17
		fmt.Println("Taxe is 17%")
		fmt.Println(income)
		fmt.Println(taxe)
	}
}

func main() {
	personIncome := 170000.00
	taxesSalary(personIncome)
}
