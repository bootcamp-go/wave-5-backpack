package main

import "fmt"

func main(){
	salario:=CalculoSalario(180,"C")
	fmt.Printf("%.2f\n",salario)
}
func CalculoSalario(minTrabajo float64, categoria string) float64 {
	var hrtrabajo,bono,salario float64 
	hrtrabajo = minTrabajo/float64(60)
	
	switch categoria {
	case "A":
		salario = hrtrabajo*3000
		
		bono = (50 * salario)/100
		fmt.Printf("hr:%.2f sal:%.2f bono:%.2f\n",hrtrabajo, salario, bono)
		salario += bono
		fmt.Printf("%.2f",salario)
	
	case "B":
		salario = hrtrabajo*1500
		bono = (20 * salario)/100
		fmt.Printf("hr:%.2f sal:%.2f bono:%.2f\n",hrtrabajo, salario, bono)
		salario += bono
		fmt.Printf("%.2f",salario)
	
	case "C":
		salario = hrtrabajo*1000
		fmt.Printf("hr:%.2f sal:%.2f\n",hrtrabajo, salario)
	}
	return salario;
}