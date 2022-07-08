package main

import "fmt"

func calc_promedio(notas ...int) float32 {
	suma := 0
	for _, value := range notas {
		if value < 0 {
			return 0
		}
		suma += value
	}
	division := float32(suma) / float32(len(notas))
	return division
}

func main() {
	promedio := calc_promedio(3, 3, 3, 3, -2)
	if promedio != 0 {
		fmt.Printf("El promedio del alumno es de: %.2f\n", promedio)
	} else {
		fmt.Println("No puede ingresar un valor negativo")
	}
}