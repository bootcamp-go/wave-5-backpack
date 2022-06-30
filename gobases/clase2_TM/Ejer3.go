package main

import "fmt"


func main(){

	var minutes int 
	var category string
	fmt.Println("inserte la cantidad de minutos trabajados:")
	fmt.Scanln(&minutes)

	fmt.Println("ingrese la categoria del trabajador A,B ó C")
	fmt.Scanln(&category)

	fmt.Printf("el salario del trabajador es = %.2f \n",calcSalary(minutes, category))

}

func calcSalary(m int , c string )float64{
	var salary float64
	switch c {
	case "A":
		salary=hours(m)*3000*1.5
	case "B":
		salary=hours(m)*1500*1.2
	case "C":
		salary=hours(m)*1000
	}
	return salary
}

func hours(minWorked int) float64{
	return float64(minWorked)/60.0
}