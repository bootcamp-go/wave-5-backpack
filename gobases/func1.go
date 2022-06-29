package main

func TaxSalary1(salary float64) float64 {
	switch {
	case salary > 150.000:
		return salary - (salary * 10 / 100)
	case salary > 50.000:
		return salary - (salary * 17 / 100)
	default:
		return salary
	}
}
