package main

import "fmt"

func main(){
    var numeroMes=5
	var mesesMap = map[int]string{1:"Enero",2:"Febrero",3:"Marzo",4:"Abril",5:"Mayo",6:"Junio",7:"Julio",8:"Agosto",9:"Setiembre",10:"Octubre",11:"Noviembre",12:"Diciembre"}
    mesesArreglo := [12]string{"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Setiembre","Octubre","Noviembre","Diciembre"}
    fmt.Println(mesesMap[numeroMes])
    fmt.Println(mesesArreglo[numeroMes-1]) 
}

/*Ejercicio 3 - A qué mes corresponde

Realizar una aplicación que contenga una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto. 
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
Ej: 7, Julio
*/