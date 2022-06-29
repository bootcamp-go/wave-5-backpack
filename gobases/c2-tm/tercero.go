package main

import "fmt"

func salario(minutos int, category string) float64 {
	var resultado float64 = 0.00
	var hora float64 = float64(minutos / 60)

	switch category {
	case "C":
		resultado = (hora * 1000.00)
	case "B":
		resultado = (hora * 1500.00) + ((hora * 1500.00) * 0.2)
	case "A":
		resultado = (hora * 3000.00) + ((hora * 3000.00) * 0.5)
	}

	return resultado
}

func main() {
	var minutos int = 12000
	var category string = "B"
	fmt.Println(salario(minutos, category))
}
