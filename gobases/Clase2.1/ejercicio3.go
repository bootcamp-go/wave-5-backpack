package main

import "fmt"

const (
	categoriaA = "A"
	categoriaB = "B"
	categoriaC = "C"
)

func calcTipo(tipo string) float64 {
	switch tipo {
	case categoriaA:
		return 3000.0
	case categoriaB:
		return 1500.0
	case categoriaC:
		return 1000.0
	}
	return 0
}

func calcAumento(tipo string) float64 {
	switch tipo {
	case categoriaA:
		return 0.5
	case categoriaB:
		return 0.2
	case categoriaC:
		return 0
	}
	return 0
}

func calcSalario(minutos int, tipo string) float64 {
	horas := float64(minutos) / 60
	neto := horas * calcTipo(tipo)
	salarioFinal := neto + neto*calcAumento(tipo)
	return salarioFinal
}

func main() {
	fmt.Println(calcSalario(60, categoriaC))
}
