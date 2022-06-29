package main

import "fmt"

func promedio(valores ...float64) float64 {
	var result float64
	count := 0.0
	for _, valor := range valores {
		count++
		result += valor
	}
	result = result / count
	return result
}

func main() {
	fmt.Println("El promedio del alumno es:", promedio(3.4, 3.5, 4.6))
}
