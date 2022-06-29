package main

import "fmt"

func main(){
	fmt.Printf("Impuesto empleado: %.2f\n",impuestoSalario(1200000))
}

func impuestoSalario(sueldo float64) float64{
	if sueldo > 50000{
		return sueldo * 0.17
	}else if(sueldo > 150000){
		return sueldo * 0.10
	}else{
		//19% de impuesto para el empleado
		//con sueldo fuera de rango
		return sueldo * 0.19
	}
}