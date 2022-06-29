package main

import "fmt"

func main(){
	fmt.Printf("Salario del trabajador es $%.2f \n",calcularSalario(120,"C"))
	fmt.Printf("Salario del trabajador es $%.2f \n",calcularSalario(120,"B"))
	fmt.Printf("Salario del trabajador es $%.2f \n",calcularSalario(120,"A"))
}

func calcularSalario(minutos int, categoria string) float64{
	//como los minutos ingresados son los trabajados por mes
	//pero importa la cantidad de horas, represento el valor de 
	//un minuto en una hora (variable minutoHora)
	salario := 0.0
	minutoHora := float64(minutos) * 0.0166667

	switch categoria{
	case "C":
		salario = 1000 * minutoHora
	case "B":
		salario = (1500 * minutoHora) + (salario * 0.2)
	case "A":
		salario = (3000 * minutoHora) + (salario * 0.5)
	}

	return salario
}