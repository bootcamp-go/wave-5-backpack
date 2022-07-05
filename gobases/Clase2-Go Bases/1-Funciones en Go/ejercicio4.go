package main

import "fmt"
import "errors"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )
 

func main(){
	//Calculo Minimo
	minFunc, err := operacionAritmetica(minimum)
	if(err!=nil){
		fmt.Printf("%v\n",err)
	}else{
		minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
		fmt.Printf("Minimo %v \n",minValue) 
	}

	//Calculo Promedio
	averageFunc, err := operacionAritmetica(average)
	if(err!=nil){
		fmt.Printf("%v\n",err)
	}else{
		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("Promedio %v \n",averageValue)  
	}
	//Calculo Maximo
	maxFunc, err := operacionAritmetica(maximum)
	if(err!=nil){
		fmt.Printf("%v\n",err)
	}else{
		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
		fmt.Printf("Maximo %v \n",maxValue)
	}

}

func operacionAritmetica(operador string) (func(valores ... int) float64,error){
	switch operador {
	case minimum:
		return minFunc,nil
	case average:
		return averageFunc,nil
	case maximum:
		return maxFunc,nil

	default:
		return nil,errors.New("Operador desconocido")
	}
	
} 

func minFunc(valores ... int) float64{
	var resultado float64=0.0
	var contador int=0
	for _,value :=range valores {
		if(contador==0||float64(value)<resultado){
			resultado=float64(value)
			contador++
		}
	}
	return resultado 
}

func averageFunc(valores ... int) float64{
	var resultado float64=0.0
	var contador=0.0
	for _,value :=range valores {
		resultado=resultado+float64(value)
		contador++
		}
	resultado=resultado/contador
	return resultado
}
func maxFunc(valores ... int) float64{
	var resultado float64=0.0
	for _,value :=range valores {
		if(float64(value)>resultado){
			resultado=float64(value)
		}
	}
	return resultado 
}





/*Ejercicio 4 - Calcular estadísticas

Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
 de calificaciones de los alumnos de un curso, requiriendo calcular los valores
  mínimo, máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
 (mínimo, máximo o promedio) y que devuelva otra función
  ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N
   de enteros y devuelva el cálculo que se indicó en la función anterior
*/