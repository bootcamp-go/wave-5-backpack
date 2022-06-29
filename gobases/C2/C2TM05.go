package main

import (
	"fmt"
)

func calcularAlimentoTarantulas(valor float64) float64 {

	return valor * 150
}

func calcularAlimentoHamsters(valor float64) float64 {

	return valor * 250
}

func calcularAlimentoPerros(valor float64) float64 {

	return valor * 10000
}

func calcularAlimentoGatos(valor float64) float64 {
	return valor * 5000
}

func errorAnimal(valor float64) float64 {
	return 0.0
}

func operacionAritmetica(operador string) func(valor float64) float64 {
	switch operador {
	case "tarantulas":
		return calcularAlimentoTarantulas
	case "hamsters":
		return calcularAlimentoHamsters
	case "perros":
		return calcularAlimentoPerros
	case "gatos":
		return calcularAlimentoGatos
	default:
		return errorAnimal
	}
}
func main() {

	tarantulas := operacionAritmetica("tarantulas")
	hamsters := operacionAritmetica("hamsters")
	perros := operacionAritmetica("perros")
	gatos := operacionAritmetica("gatos")

	AlimentoTarantulas := tarantulas(1.0)
	AlimentoHamsters := hamsters(4.0)
	AlimentoPerros := perros(5.0)
	AlimentoGatos := gatos(3.0)

	fmt.Println("Alimento para tarantulas: ", AlimentoTarantulas, "gramos o", AlimentoTarantulas/1000, "Kg")
	fmt.Println("Alimento para hamsters: ", AlimentoHamsters, "gramos o ", AlimentoHamsters/1000, "Kg")
	fmt.Println("Alimento para perros: ", AlimentoPerros, "gramos o ", AlimentoPerros/1000, "Kg")
	fmt.Println("Alimento para gatos: ", AlimentoGatos, "gramos o ", AlimentoGatos/1000, "Kg")

}
