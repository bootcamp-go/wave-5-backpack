package main

import (
	"fmt"
)

func calcularSalarioPorHora(categoria string)  (salarioPorHora float64){
	switch categoria {
	case "A":
		salarioPorHora=3000.0
	case "B":
		salarioPorHora=1500.0
	case "C":
		salarioPorHora=1000.0
	default:
		salarioPorHora=0.0
	}
	return
}

func calcularSalario(minutos float64, categoria string)  float64{
	horas := minutos/60
	salarioPorHora := calcularSalarioPorHora(categoria)

	return horas * salarioPorHora
}
func main() {

	fmt.Println(calcularSalario(120,"C"))

}