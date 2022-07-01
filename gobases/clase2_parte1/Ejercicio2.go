package main

import "fmt"

func average(values ...int) int {
	var resultado int
	for _, value := range values {
		resultado += value
	}
	return resultado / len(values)
}

func main() {
	fmt.Println("El promedio es :", average(10, 10, 10, 11, 12))
}
