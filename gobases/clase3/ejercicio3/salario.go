package main

import "fmt"

func salarioCategoriaA(tiempoTrabajoH float64) float64 {
	salario := tiempoTrabajoH * 3000
	return salario + (salario * 0.5)
}
func salarioCategoriaB(tiempoTrabajoH float64) float64 {
	salario := tiempoTrabajoH * 1500
	return salario + (salario * 0.2)
}
func salarioCategoriaC(tiempoTrabajoH float64) float64 {
	return tiempoTrabajoH * 1000
}

func calculoSalario(tiempoTrabajoMin float64, categoria string) float64 {
	tiempoTrabajoH := tiempoTrabajoMin / 60
	fmt.Println("Categoria:", categoria)
	fmt.Println("Tiempo de trabajo en horas:", tiempoTrabajoH)

	switch categoria {
	case "A":
		return salarioCategoriaA(tiempoTrabajoH)
	case "B":
		return salarioCategoriaB(tiempoTrabajoH)
	case "C":
		return salarioCategoriaC(tiempoTrabajoH)

	}

	return 0

}

func main() {
	fmt.Println("El salario total del trabajador es:", calculoSalario(80, "A"))
}
