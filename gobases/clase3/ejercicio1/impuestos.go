package main

import "fmt"

func taxResult(salario float64) float64 {
	var result float64

	if salario > 50000 {
		result = salario - (salario * 0.17)
	} else if salario > 150000 {
		result = salario - (salario * 0.10)
	}
	return result
}

func main() {

	fmt.Println("Tu sueldo total es:", taxResult(160000))
}
