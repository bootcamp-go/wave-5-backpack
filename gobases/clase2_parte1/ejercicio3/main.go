package main

import "fmt"

func sueldo(operador string, minutos int) float64 {
	a := float64(minutos)
	horas := a / 60
	switch operador {
	case "C":
		return horas * 1000
	case "B":
		sueldoMensual := horas * 1500
		return sueldoMensual * 1.20
	case "A":
		sueldoMensual := horas * 3000
		return sueldoMensual * 1.50
	}
	return 0

}

func main() {

	res := sueldo("C", 600)

	if res != 0 {
		fmt.Println("Sueldo:", res)

	} else {
		fmt.Println("No existe la categoria")
	}

}
