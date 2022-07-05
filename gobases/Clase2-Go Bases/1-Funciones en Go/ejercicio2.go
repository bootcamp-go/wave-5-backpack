package main

import "fmt"
import "errors"

func main(){
	promedio,falla:=calcularPromedio(-12,11,5,7,8,9,10)
	if(falla!=nil){
		fmt.Printf("%v",falla)
	}else{
		fmt.Printf("El promedio es %v \n",promedio)
	}
	 
}

func calcularPromedio(valores ... float64) (float64,error){
	var resultado float64=0.0
	var contador=0.0
	for _,value :=range valores {
		if(value<0){
			return 0,errors.New("Existen notas negativas\n") 
		}else{
			resultado=resultado+value
		contador++
		}
		
	}
	resultado=resultado/contador
	return resultado,nil
} 





/*
Ejercicio 2 - Calcular promedio

Un colegio necesita calcular el promedio (por alumno) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y 
devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

*/