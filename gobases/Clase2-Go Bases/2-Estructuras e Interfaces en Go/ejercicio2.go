package main

import "fmt"

type Matrix struct {
	Valores []float64
	Alto int
	Ancho int
	Cuadratica bool
	ValorMaximo float64
}
func (matriz *Matrix) Set(valores ... float64){
	matriz.Valores =valores
	matriz.Cuadratica= matriz.Alto==matriz.Ancho
	matriz.Alto= 
	//matriz.ValorMaximo=maxFunc(valores)
}

func maxFunc(valores ... float64) float64{
	var resultado float64=0.0
	for _,value :=range valores {
		if(value>resultado){
			resultado=value
		}
	}
	return resultado 
}

//func (matriz Matrix) Print(){
//	salida:="Nombre: "+individuo.Nombre+"\nApellido: "+individuo.Apellido+"\nDNI: "+individuo.DNI+"\nFecha: "+individuo.FechaIngreso+"\n"
//	fmt.Printf(salida) 
//}
func main(){

	matriz:=Matrix{}
	matriz.Set(1.2,3.67,343,1005.3)
	fmt.Print(matriz)
}

/*
Ejercicio 2 - Matrix
Una empresa de inteligencia artificial necesita tener una funcionalidad para crear 
una estructura que represente una matriz de datos.
Para ello requieren una estructura Matrix que tenga los métodos:
Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, 
la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

*/