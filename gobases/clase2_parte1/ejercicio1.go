package main

import "fmt"

func main() {
	var sueldos = []float64{20000, 30000, 55000, 98000, 160000}
	fmt.Println(calcTaxes(sueldos))
}

func calcTaxes(sueldos []float64) []float64 {
	for idx, sueldo := range sueldos {
		if sueldo > 50000 && sueldo < 160000 {
			sueldos[idx] = sueldo * 0.83
		} else if sueldo > 150000 {
			sueldos[idx] = sueldo * 0.73
		}
	}
	return sueldos
}
