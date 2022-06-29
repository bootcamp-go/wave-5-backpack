package main

import (
	"fmt"
	"math"
)

func main() {
	meanPromedio(3.1, 2.33333, 4.35, 4.3, 3)
	meanPromedio(3.1, 2.33333, -4.35, 4.3, 3)
}

func meanPromedio(notas ...float64) {
	var promedio float64 = 0
	for _, nota := range notas {
		if nota < 0 {
			fmt.Println("Los valores ingresados no pueden ser negativos")
			return
		}
		promedio += float64(nota)
	}
	promedio = promedio / float64(len(notas))
	promedio = math.Round(promedio*100) / 100
	fmt.Printf("El promedio de las %d notas es de: %f\n", len(notas), promedio)
}
