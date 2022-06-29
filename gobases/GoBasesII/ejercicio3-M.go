package main

import "fmt"

func main() {
	oper := mTrabajados("C")
	r := oper(1920)

	fmt.Println(r)
}

func A(min float64) float64 {
	var horas float64 = min / 360
	return 3000 * horas * (1.50)
}

func B(min float64) float64 {
	var horas float64 = min / 360
	return 1500 * horas * (1.20)
}

func C(min float64) float64 {
	var horas float64 = min / 360
	return 1000 * horas
}

func mTrabajados(categoria string) func(min float64) float64 {

	switch categoria {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	}
	return nil
}
