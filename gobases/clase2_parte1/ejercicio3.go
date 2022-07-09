package main

import "fmt"

func main()  {
	fmt.Printf("El salario es de %.2f\n",CalculoSalario(480, "C"))
}

func CalculoSalario(minutos int , categoria string)  float64 {

	horas := minutos / 60
	switch categoria {
	case "C":
		return float64(horas*1000)
	case "B":
		salario := horas * 1500
		bonus := float64(salario) * float64(0.20)
		resultado := float64(salario) + float64(bonus)
		return resultado
	case "A":
		salario := horas * 3000
		bonus := float64(salario) * float64(0.50)
		resultado := float64(salario) + float64(bonus)
		return resultado
	default:
		return float64(0)
	}
}
