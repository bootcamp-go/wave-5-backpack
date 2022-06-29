package main

import "fmt"

func main(){
	var meses = map[int]string{1:"ENERO" , 2:"FEBRERO", 3:"MARZO", 4:"ABRIL",5:"MAYO",6:"JUNIO",7:"JULIO",8:"AGOSTO",9:"SETIEMBRE",10:"OCTUBRE",11:"NOVIEMBRE",12:"DICIEMBRE" }
		var rec int = 11
		if (rec>0 && rec <12){
			fmt.Println("El mes es :",meses[rec])
		} else { 
			fmt.Println("El número ingresado no es correcto")
		}
		
}